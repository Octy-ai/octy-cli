package rest

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	cp "github.com/Octy-ai/octy-cli/internal/application/domain/churn_prediction_report"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

func (ha Adapter) GetChurnReport(credentials string) (*cp.OctyChurnPredictionReport, []error) {

	req, err := http.NewRequest("GET", globals.GetChurnReportRoute, nil)
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

	getChurnPredictionReportResp, err := models.UnmarshalOctyGetChurnPredictionReportResp(body)
	if err != nil {
		return nil, []error{err}
	}

	churnPredictionReport := cp.NewCPR()
	copier.Copy(&churnPredictionReport.ChurnPredictionReport, &getChurnPredictionReportResp.ChurnPredictionReport)

	return churnPredictionReport, nil
}
