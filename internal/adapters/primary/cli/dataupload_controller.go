package cli

import (
	"fmt"

	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	"github.com/Octy-ai/octy-cli/pkg/output"
)

func uploadDataController(clia Adapter, data *d.Data) {
	//spinner := output.StartNewSpinner(fmt.Sprintf("validating %s csv file ...", data.ResourceType), os.Stdout)
	content, objectIDXMap, errs := clia.api.ValidateData(data)
	if errs != nil {
		//errorFactory(errs, true, spinner)
		errorFactory(errs, true, nil)
	}
	// kill spinner if validations pass
	//spinner.FinalMSG = fmt.Sprintf("> Valid file. Processing %s file upload \n\n", data.ResourceType)
	//spinner.Stop()

	// create channel to track upload progress
	progressChan := make(chan d.UploadProgess, 100)

	errs = clia.api.UploadData(data.ResourceType, objectIDXMap, content, progressChan)
	if errs != nil {
		errorFactory(errs, true, nil)
	}

	// listen on channel
	d.Wg.Add(1)
	go func(resourceType string, pch <-chan d.UploadProgess) {

		var report = map[string]interface{}{
			"completeCount":  0,
			"completeChunks": 0,
			"failedCount":    0,
			"failed":         []d.Failed{},
		}

	ChannelLoop:
		for {
			if i, ok := <-pch; ok {
				// update progress report
				complete := report["completeCount"].(int)
				complete += int(i.Complete)
				report["completeCount"] = complete
				report["failedCount"] = int(len(i.Failed))
				report["failed"] = i.Failed
			} else {
				// output upload report.
				fmt.Println()
				fmt.Println("--")
				fmt.Println()
				fmt.Println("Complete. All parts processed.")
				fmt.Println()
				fmt.Println("Report:")
				fmt.Printf("* Created a total of %v %v \n", report["completeCount"], resourceType)
				fmt.Printf("* Failed to create %v %v\n \n", report["failedCount"], resourceType)
				// output error report if errors

				if report["failedCount"].(int) > 0 {
					fmt.Println("Error Report:")
					errOutCount := 0
					for _, e := range report["failed"].([]d.Failed) {
						if errOutCount >= 10 {
							output.FPrint(fmt.Sprintf("Too many Errors to log. %v more errors found ...", len(report["failed"].([]d.Failed))-10))
							break ChannelLoop // exit loop and stop listening on channel
						}
						output.FPrint("* Error:")
						fmt.Printf(" Error Message: %v \n", e.ErrorMessage)
						fmt.Printf(" Row Index: %v \n", e.RowIDX)
						errOutCount++
					}
				}
				break ChannelLoop // exit loop and stop listening on channel
			}
		}
		d.Wg.Done()

	}(data.ResourceType, progressChan)

	d.Wg.Wait()
}
