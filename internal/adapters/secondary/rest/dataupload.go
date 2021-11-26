package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/utils"
)

func (ha Adapter) GetResourceFormats(resourceType string) (map[string]string, map[int]string, error) {

	req, err := http.NewRequest("GET", globals.GetResourceFormats+"?type="+resourceType, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	switch {
	case resp.StatusCode > 200 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, models.ParseErrors(errResp)[0]
	case resp.StatusCode >= 500:
		return nil, nil, errors.New("apierror[500]:: unknown server error")
	}

	var respBody map[string]interface{}
	var expected = make(map[string]string)
	var expectedKeys = make(map[int]string)

	json.Unmarshal(body, &respBody)

	for i, v := range respBody {
		columnName := utils.BeforeStr(i, "**")
		columnIDX, _ := strconv.Atoi(utils.AfterStr(i, "**"))
		expectedKeys[columnIDX] = columnName

		switch v := v.(type) {
		case string:
			if v == "nested" {
				expected[columnName] = "nested"
				continue
			}
			expected[columnName] = "string"
		case float64:
			//floats should be provided in ref response as:
			// 0, int in ref as: 1
			if v < 1 {
				expected[columnName] = "float"
				continue
			}
			expected[columnName] = "int"
		case bool:
			expected[columnName] = "bool"
		}
	}

	return expected, expectedKeys, nil
}

func (ha Adapter) UploadProfiles(profiles string, objectRowIDXMap *map[string]int, credentials string, prog *d.UploadProgess, progressChan chan<- d.UploadProgess) {

	defer d.Wg.Done()

	var errs []error
	bodyJSON := "{ \"profiles\" : " + profiles + "}"

	req, err := http.NewRequest("POST", globals.CreateProfiles, bytes.NewBuffer([]byte(bodyJSON)))
	if err != nil {
		errs = append(errs, err)
	}
	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")
	resp, err := ha.httpClient.Do(req)
	if err != nil {
		errs = append(errs, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errs = append(errs, err)
	}

	switch {
	case resp.StatusCode > 201 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			errs = append(errs, err)
		}
		errs = append(errs, models.ParseErrors(errResp)...)
	case resp.StatusCode >= 500:
		errs = append(errs, errors.New("apierror[500]:: unknown server error"))
	}

	createProfilesResp, err := models.UnmarshalOctyCreateProfilesResp(body)
	if err != nil {
		errs = append(errs, err)
	}

	// update prog values and push to channel
	prog.Complete = int64(len(createProfilesResp.Profiles))
	for _, f := range createProfilesResp.FailedToCreate {

		failed := d.NewFailed()
		failed.ErrorMessage = f.ErrorMessage
		failed.RowIDX = (*objectRowIDXMap)[f.CustomerID]
		prog.Failed = append(prog.Failed, *failed)

	}
	for _, e := range errs {
		failed := d.NewFailed()
		failed.ErrorMessage = e.Error()
		if strings.Contains(e.Error(), "customer_id :") {
			failed.RowIDX = (*objectRowIDXMap)[utils.AfterStr(e.Error(), "customer_id : ")]
		} else {
			failed.RowIDX = 0
		}
		prog.Failed = append(prog.Failed, *failed)
	}

	prog.Bar.Add(1) // bump progess bar
	progressChan <- *prog
	prog.Mutex.Unlock()

}
