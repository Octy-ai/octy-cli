package rest

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	v "github.com/Octy-ai/octy-cli/internal/application/domain/versioning"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

func (ha Adapter) GetVersionInfo() (*v.Version, error) {

	req, err := http.NewRequest("GET", globals.CLIVersionInfo, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode > 200 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, err
		}
		return nil, models.ParseErrors(errResp)[0]
	case resp.StatusCode >= 500:
		return nil, errors.New("apierror[500]:: unknown server error")
	}

	getVersionInfoResp, err := models.UnmarshalOctyGetVersionInfoResp(body)
	if err != nil {
		return nil, err
	}

	version := v.NewVersion()
	copier.Copy(&version, &getVersionInfoResp.CurrentVersion)

	return version, nil
}
