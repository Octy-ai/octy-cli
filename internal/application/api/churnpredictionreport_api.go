package api

import cp "github.com/Octy-ai/octy-cli/internal/application/domain/churn_prediction_report"

//
// Public methods
//

func (api Application) GetChurnReport() (*cp.OctyChurnPredictionReport, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetChurnReport(credentials)
}
