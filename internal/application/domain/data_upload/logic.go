package data_upload

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/utils"
)

//
// Public upload methods
//

// GetReferenceMaps : Retuns map containing provided column headers, a map to track the indices of each provided column header and a map to track the row index of each unique identifier.
func (u *Upload) GetReferenceMaps(content *[][]string, resourceType string) (map[string]string, map[int]string, *map[string]int, []error) {

	var columns = make([]string, 0)
	columns = append(columns, (*content)[0]...)
	var provided = make(map[string]string)
	var providedKeys = make(map[int]string)
	objectRowIDXMap := map[string]int{}
	duplicates := make(map[string][]int)
	var identifier string

	switch resourceType {
	case "profiles":
		identifier = "customer_id"
	}

RowLoop:
	for rowIDX, row := range *content {
		if rowIDX == 0 {
			keyIdx := 0
			for colIDX, colVal := range columns {
				val := colVal
				if strings.Contains(colVal, ">") {
					val = utils.BeforeStr(colVal, ">")
				}
				if !utils.ValueInIntStrMap(val, providedKeys) {
					providedKeys[colIDX] = val
					keyIdx++
					continue
				}
			}
			continue
		}
		row = trimCellWhiteSpace(row)
		for colIDX, colVal := range columns {

			if resourceType != "events" { // can not track indices of events. No unique idenitifers.

				// populate object index map
				if colVal == identifier {

					// check for duplicates (items / profiles only)
					if _, found := objectRowIDXMap[row[colIDX]]; found {
						duplicates[row[colIDX]] = append(duplicates[row[colIDX]], rowIDX)
						continue RowLoop
					}
					duplicates[row[colIDX]] = append(duplicates[row[colIDX]], rowIDX)
					objectRowIDXMap[row[colIDX]] = rowIDX
				}
			}

			cellType := "string"
			if globals.BoolRepresentations[row[colIDX]] {
				cellType = "bool"
			} else {

				if strings.Contains(row[colIDX], ".") {
					if _, err := strconv.ParseFloat(row[colIDX], 64); err == nil {
						cellType = "float"
					}
				} else {
					if _, err := strconv.ParseUint(row[colIDX], 10, 32); err == nil {
						cellType = "int"
					}
				}
			}
			_, ok := provided[colVal]
			if ok {
				colType := provided[colVal]
				if colType != cellType {
					return nil, nil, nil, []error{fmt.Errorf("validationerror[invalid]:: mismatched data types provided in column %q , at row index %v", colVal, rowIDX+1)}
				}
			} else {
				provided[colVal] = cellType
			}
		}
	}

	// check for duplicate errors
	var errs []error
	for k, v := range duplicates {
		if len(v) > 1 {
			errs = append(errs, fmt.Errorf("validationerror[duplicate]:: duplicate object identifier %q found at row indices %v ", k, v))
		}
	}
	if len(errs) > 0 {
		return nil, nil, nil, errs
	}

	parsedProvided := make(map[string]string)
	for i, v := range provided {
		if strings.Contains(i, ">") {
			parsedProvided[utils.BeforeStr(i, ">")] = "nested"
			continue
		}
		parsedProvided[i] = v
	}
	return parsedProvided, providedKeys, &objectRowIDXMap, nil
}

// Profiles --

// ParseProfiles : parse profiles data to required types and shape
func (u *Upload) ParseProfiles(content *[][]string) (*[][]string, []string) {

	// init empty 2d array with length of content minus the header row
	contentC := make([][]string, len(*content)-1)

	// get column names
	columns := make([]string, 0)
	columns = append(columns, (*content)[0]...)
	//row index, row value
	for rowIDX, rowValue := range (*content)[1:] {
		rowValue = trimCellWhiteSpace(rowValue)
		profileDataMap := make(map[string]interface{})
		platformInfoMap := make(map[string]interface{})

		// iterate over headers -> representing row columns headers
		for colIDX, colValue := range columns {
			if !strings.Contains(colValue, ">") {
				contentC[rowIDX] = append(contentC[rowIDX], rowValue[colIDX])
				continue
			}
			t := toRepsentedType(rowValue[colIDX])
			switch utils.BeforeStr(colValue, ">") {
			case "profile_data":
				switch v := t.(type) {
				case string:
					profileDataMap[utils.AfterStr(colValue, ">")] = v
				case int:
					profileDataMap[utils.AfterStr(colValue, ">")] = v
				case float64:
					profileDataMap[utils.AfterStr(colValue, ">")] = v
				case bool:
					profileDataMap[utils.AfterStr(colValue, ">")] = v
				}
			case "platform_info":
				switch v := t.(type) {
				case string:
					platformInfoMap[utils.AfterStr(colValue, ">")] = v
				case int:
					platformInfoMap[utils.AfterStr(colValue, ">")] = v
				case float64:
					platformInfoMap[utils.AfterStr(colValue, ">")] = v
				case bool:
					platformInfoMap[utils.AfterStr(colValue, ">")] = v
				}
			}
		}

		profileDataJSON, err := json.Marshal(profileDataMap)
		if err != nil {
			continue
		}
		platformInfoJSON, err := json.Marshal(platformInfoMap)
		if err != nil {
			continue
		}
		contentC[rowIDX] = append(contentC[rowIDX], string(profileDataJSON), string(platformInfoJSON))
	}

	return &contentC, columns
}

// Events --

// ParseEvents : parse events data to required types and shape
// func (u *Upload)  ParseEvents(content *[][]string) *[][]string {

// }

//
// Private upload functions
//

// trimCellWhiteSpace : remove the whitespace from each cell in the provided row.
func trimCellWhiteSpace(row []string) []string {
	trimmedRow := []string{}
	for _, v := range row {
		trimmedRow = append(trimmedRow, strings.TrimSpace(v))
	}
	return trimmedRow
}

// toRepsentedType : convert string to string representation of type
func toRepsentedType(cellValue string) interface{} {

	if globals.BoolRepresentations[cellValue] {
		b, err := strconv.ParseBool(cellValue)
		if err == nil {
			return b
		}
	}
	if globals.NullRepresentations[cellValue] {
		return "null"
	}
	if strings.Contains(cellValue, ".") {
		if f, err := strconv.ParseFloat(cellValue, 64); err == nil {
			return f
		}
	} else {
		i, err := strconv.ParseUint(cellValue, 10, 64)
		if err == nil {
			return int(i)
		}
	}

	return cellValue
}

// ---

//
// Semaphored wait group that implements the WaitGroup interface.
//

type SemaphoredWaitGroup struct {
	sem chan bool
	wg  sync.WaitGroup
}

func (s *SemaphoredWaitGroup) Add(delta int) {
	s.wg.Add(delta)
	s.sem <- true
}

func (s *SemaphoredWaitGroup) Done() {
	<-s.sem
	s.wg.Done()
}

func (s *SemaphoredWaitGroup) Wait() {
	s.wg.Wait()
}

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

// globally accessible semaphored waitgroup
var Wg = SemaphoredWaitGroup{sem: make(chan bool, 10)}
