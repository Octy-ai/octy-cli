package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

// Account configurations set get

func (ha Adapter) SetAccountConfigurations(octyAccConfig *c.OctyAccountConfigurations, credentials string) []error {
	setAccConfigReq := models.OctySetAccConfigReq{
		ContactName:         octyAccConfig.ContactName,
		ContactSurname:      octyAccConfig.ContactSurname,
		ContactEmailAddress: octyAccConfig.ContactEmailAddress,
		WebhookURL:          octyAccConfig.WebhookURL,
		AuthenticatedIDKey:  octyAccConfig.AuthenticatedIDKey,
	}

	requestBody, err := setAccConfigReq.Marshal()
	if err != nil {
		return []error{err}
	}
	req, err := http.NewRequest("POST", globals.SetAccConfigRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		return []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []error{err}
	}

	switch {
	case resp.StatusCode > 202 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return []error{err}
		}
		return models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return []error{errors.New("apierror[500]:: unknown server error")}
	}

	return nil
}

func (ha Adapter) GetAccountConfigurations(credentials string) (*c.OctyAccountConfigurations, []error) {

	req, err := http.NewRequest("GET", globals.GetAccConfigRoute, nil)
	if err != nil {
		return nil, []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, []error{err}
	}

	switch {
	case resp.StatusCode > 200 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, []error{err}
		}
		return nil, models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return nil, []error{errors.New("apierror[500]:: unknown server error")}
	}

	getAccConfigResp, err := models.UnmarshalOctyGetAccConfigResp(body)
	if err != nil {
		return nil, []error{err}
	}

	accountConf := c.NewAcc()
	copier.Copy(&accountConf, &getAccConfigResp.AccountData)

	return accountConf, nil
}

// Algorithm configurations set get

func (ha Adapter) SetRecAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) []error {

	setRecAlgoConfigReq := models.OctySetRecAlgoConfigReq{
		AlgorithmName: octyAlgoConfig.AlgorithmName,
		Configurations: models.RecReqConfigurations{
			RecommendInteractedItems: octyAlgoConfig.Configurations.RecommendInteractedItems,
			ItemIDStopList:           octyAlgoConfig.Configurations.ItemIDStopList,
			ProfileFeatures:          octyAlgoConfig.Configurations.ProfileFeatures,
		},
	}

	requestBody, err := setRecAlgoConfigReq.Marshal()
	if err != nil {
		return []error{err}
	}
	req, err := http.NewRequest("POST", globals.SetAlgoConfigRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		return []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []error{err}
	}

	switch {
	case resp.StatusCode > 202 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return []error{err}
		}
		return models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return []error{errors.New("apierror[500]:: unknown server error")}
	}

	return nil
}

func (ha Adapter) SetChurnAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) []error {

	setChurnAlgoConfigReq := models.OctySetChurnAlgoConfigReq{
		AlgorithmName: octyAlgoConfig.AlgorithmName,
		Configurations: models.ChurnReqConfigurations{
			ProfileFeatures: octyAlgoConfig.Configurations.ProfileFeatures,
		},
	}

	requestBody, err := setChurnAlgoConfigReq.Marshal()
	if err != nil {
		return []error{err}
	}
	req, err := http.NewRequest("POST", globals.SetAlgoConfigRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		return []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []error{err}
	}

	switch {
	case resp.StatusCode > 202 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return []error{err}
		}
		return models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return []error{errors.New("apierror[500]:: unknown server error")}
	}

	return nil
}

func (ha Adapter) GetAlgorithmConfigurations(credentials string) (*[]c.OctyAlgorithmConfiguration, []error) {

	req, err := http.NewRequest("GET", globals.GetAlgoConfigRoute, nil)
	if err != nil {
		return nil, []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, []error{err}
	}

	switch {
	case resp.StatusCode > 200 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, []error{err}
		}
		return nil, models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return nil, []error{errors.New("apierror[500]:: unknown server error")}
	}

	getAlgoConfigResp, err := models.UnmarshalOctyGetAlgoConfigResp(body)
	if err != nil {
		return nil, []error{err}
	}

	algoConf := c.NewAlg()
	copier.Copy(&algoConf, &getAlgoConfigResp.Configurations)

	return algoConf, nil
}
