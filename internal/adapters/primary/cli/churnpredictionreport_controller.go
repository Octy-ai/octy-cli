package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	c "github.com/Octy-ai/octy-cli/internal/application/domain/churn_prediction_report"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/output"
	"github.com/olekukonko/tablewriter"
)

func getChurnPredictionReportController(clia Adapter, outpath string) {
	spinner := output.StartNewSpinner("Getting churn prediction report...", os.Stdout)
	cpr, errs := clia.api.GetChurnReport()
	if errs != nil {
		errorFactory(errs, true, spinner)
	}
	if outpath != "" {
		filepath := outpath + "octy-churn-report--" + time.Now().Format("2006-01-02 15:04:05") + ".MD"
		renderChurnReportMD(cpr, filepath)
		quit(fmt.Sprintf("Octy churn prediction report created at : %v \n", filepath), 0, spinner)
	} else {
		renderChurnReport(cpr)
		quit("", 0, spinner)
	}
}

//
// Private functions
//

func renderChurnReport(cp *c.OctyChurnPredictionReport) {

	fmt.Println("--")
	fmt.Println("Churn prediction report :: ")

	// init new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Report section", "Values"})

	//init empty table data arr
	tableData := [][]string{}
	tableData = append(tableData, []string{"Training job data", "Training job ID : " + cp.ChurnPredictionReport.TrainingJobData.TrainingJobID})
	tableData = append(tableData, []string{"Training job data", "Model accuracy : " + strconv.FormatFloat(cp.ChurnPredictionReport.TrainingJobData.ModelAccuracy, 'f', 6, 64)})
	tableData = append(tableData, []string{"Training job data", "Training job date : " + cp.ChurnPredictionReport.TrainingJobData.TrainingJobDate})

	tableData = append(tableData, []string{"Churn data", "Current churn % : " + strconv.FormatFloat(cp.ChurnPredictionReport.ChurnData.CurrentChurnPercentage, 'f', 6, 64)})
	tableData = append(tableData, []string{"Churn data", "Direction indication : " + cp.ChurnPredictionReport.ChurnData.ChurnDirectionIndication})
	tableData = append(tableData, []string{"Churn data", "Percentage difference : " + strconv.FormatFloat(cp.ChurnPredictionReport.ChurnData.ChurnPercentageDifference, 'f', 6, 64)})

	for _, v := range cp.ChurnPredictionReport.ChurnData.FeaturesOfImportance {

		s := strconv.FormatFloat(v.FeatureImportance, 'f', 6, 64)
		tableData = append(tableData, []string{"Features of importance:", v.FeatureName + " : " + s})

	}

	// populate table and render
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(tableData)
	table.Render()

}

func renderChurnReportMD(cp *c.OctyChurnPredictionReport, filepath string) {

	rawMD := (`
# Octy Churn prediction report -- ` + time.Now().Format("2006-01-02 15:04:05") + `

## Training job data
	
**Training job id :** ` + cp.ChurnPredictionReport.TrainingJobData.TrainingJobID + `
	
**Model accuracy :** ` + strconv.FormatFloat(cp.ChurnPredictionReport.TrainingJobData.ModelAccuracy, 'f', 3, 64) + `
	
**Training job date :** ` + cp.ChurnPredictionReport.TrainingJobData.TrainingJobDate + `
	
## Churn data
	
**Current churn percentage :** ` + strconv.FormatFloat(cp.ChurnPredictionReport.ChurnData.CurrentChurnPercentage, 'f', 2, 64) + `%
	
**Direction indication :** ` + cp.ChurnPredictionReport.ChurnData.ChurnDirectionIndication + `
	
**Percentage difference :** ` + strconv.FormatFloat(cp.ChurnPredictionReport.ChurnData.ChurnPercentageDifference, 'f', 2, 64) + `%
	
## Features of importance
`)

	for _, v := range cp.ChurnPredictionReport.ChurnData.FeaturesOfImportance {

		rawMD += "- " + v.FeatureName + " : " + strconv.FormatFloat(v.FeatureImportance, 'f', 3, 64) + "\n"

	}

	// build footer
	rawMD += "\n *Octy churn prediction report automatically generated using the octy-cli v." + globals.CliVersion + "*"

	//footer image
	rawMD += "     <img src=\"" + globals.OctyLogoURL + "\" alt=\"drawing\" width=\"100\"/>"

	mdBytes := []byte(rawMD)

	ioutil.WriteFile(filepath, mdBytes, 0644)
}
