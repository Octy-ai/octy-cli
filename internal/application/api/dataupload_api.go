package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

//
// Public methods
//

// ValidateData : perform validation checks on provided data set.
func (api Application) ValidateData(data *d.Data) (*[][]string, *map[string]int, []error) {

	// verfiy the provided column headers match expected column headers. Both in order and names
	expected, expectedKeys, err := api.rest.GetResourceFormats(data.ResourceType)
	if err != nil {
		return nil, nil, []error{err}
	}
	// read file into memory
	reader := csv.NewReader(bytes.NewBuffer(*data.Data))
	reader.TrimLeadingSpace = true
	content, err := reader.ReadAll()
	if err != nil {
		return nil, nil, []error{err}
	}
	// verify total number of data rows
	if len(content)-1 > globals.UploadLimit {
		return nil, nil, []error{fmt.Errorf("validationerror[limit]:: Limit of %v rows exceeded", globals.UploadLimit)}
	}

	provided, providedKeys, objectRowIDXMap, errs := api.upload.GetReferenceMaps(&content, data.ResourceType)
	if errs != nil {
		return nil, nil, errs
	}

	errs = validateColumnHeaders(provided, providedKeys, expected, expectedKeys)
	if errs != nil {
		return nil, nil, errs
	}

	return &content, objectRowIDXMap, nil
}

// UploadData : upload data to Octy services.
func (api Application) UploadData(resourceType string, objectRowIDXMap *map[string]int, content *[][]string, progressChan chan<- d.UploadProgess) []error {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return []error{err}
	}

	var c *[][]string
	var columns []string
	var identifier string

	switch resourceType {
	case "profiles":
		c, columns = api.upload.ParseProfiles(content)
		identifier = "customer_id"
	case "items":
		c, columns = api.upload.ParseItems(content)
		identifier = "item_id"
	}

	// chunk upload data
	jsonBodySlice, err := chunk(identifier, c, columns)
	if err != nil {
		return []error{err}
	}
	totalChunks := int64(len(*jsonBodySlice))

	prog := d.NewUploadProgess()
	prog.Total = int64(len(*c) - 1)
	prog.TotalChunks = totalChunks

	bar := progressbar.NewOptions(int(totalChunks),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(45),
		progressbar.OptionSetDescription(fmt.Sprintf("[reset] Uploading %s data... ", resourceType)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	prog.Bar = bar

	for _, chunk := range *jsonBodySlice { // init goroutine for each chunk of data
		d.Wg.Add(1)
		prog.Mutex.Lock() // set mutex write lock before spawning each goroutine
		switch resourceType {
		case "profiles":
			go api.rest.UploadProfiles(chunk, objectRowIDXMap, credentials, prog, progressChan)
		case "items":
			go api.rest.UploadItems(chunk, objectRowIDXMap, credentials, prog, progressChan)
		}
	}
	d.Wg.Wait()
	close(progressChan)

	return nil
}

//
// Private functions
//

// validateColumnHeaders : validate provided columns are in the correct order and have the correct names/ data types.
func validateColumnHeaders(provided map[string]string, providedKeys map[int]string, expected map[string]string, expectedKeys map[int]string) []error {

	// valdate header names and data types
	if !cmp.Equal(expected, provided) {
		return []error{fmt.Errorf("validationerror[invalid]:: invalid csv file columns headers or data types provided. Mismatch (-want +got):\n%s", cmp.Diff(expected, provided))}
	}

	var errs []error

	// validate column header order
	for k, v := range expectedKeys {
		if v != providedKeys[k] {
			errs = append(errs, fmt.Errorf("validationerror[invalid]:: invalid column header order. mismatch (-want +got): \n - %v \n + %v \n ", expectedKeys, providedKeys))
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

// chunk : returns slice of chunked data
func chunk(identifier string, content *[][]string, columns []string) (*[]string, error) {
	var jsonBodySlice []string
	chunkCount := math.Ceil(float64(len(*content)) / float64(500)) // 500 objects per chunk
	start := 0
	for i := 0; i < int(chunkCount); i++ {
		end := math.Min(float64(len(*content)), float64(start+500))
		rows := (*content)[start:int(end)]
		rawJSON, err := sliceToJSON(identifier, columns, &rows)
		jsonBodySlice = append(jsonBodySlice, string(rawJSON))
		if err != nil {
			return nil, err
		}
		start = int(end)
	}
	return &jsonBodySlice, nil
}

// sliceToJSON : converts a slice of a slice of data into a JSON string
func sliceToJSON(identifier string, columns []string, content *[][]string) ([]byte, error) {

	columnHeaders := []string{}
	// parse column headers
	for _, column := range columns {
		val := column
		if strings.Contains(column, ">>") {
			val = utils.BeforeStr(column, ">")
		}
		res := utils.InSlice(val, columnHeaders)
		if !res {
			columnHeaders = append(columnHeaders, val)
		}
	}

	// build string buffer from csv data
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range *content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + columnHeaders[j] + `":`)
			if columnHeaders[j] != identifier {
				_, fErr := strconv.ParseFloat(y, 32)
				_, bErr := strconv.ParseBool(y)
				if fErr == nil {
					buffer.WriteString(y)
				} else if bErr == nil {
					buffer.WriteString(strings.ToLower(y))
				} else {
					if strings.Contains(y, "{") {
						buffer.WriteString((y))
					} else {
						buffer.WriteString((`"` + y + `"`))
					}
				}
			} else {
				// Identifiers should be parsed to string always, regardless of inferred type
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}
		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(*content)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(`]`)

	// convert string buffer to raw json
	rawMessage := json.RawMessage(buffer.String())
	js, err := json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		return []byte{}, err
	}

	// return json
	return js, nil
}
