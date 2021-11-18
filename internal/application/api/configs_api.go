package api

import c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"

//
// Public methods
//

func (api Application) SetAccountConfigs(octyAccConfig *c.OctyAccountConfigurations) []error {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return []error{err}
	}
	return api.rest.SetAccountConfigurations(octyAccConfig, credentials)
}

func (api Application) GetAccountConfigs() (*c.OctyAccountConfigurations, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetAccountConfigurations(credentials)
}

func (api Application) SetAlgorithmConfigs(octyAlgoConfigs *[]c.OctyAlgorithmConfiguration) []error {

	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return []error{err}
	}

	var errs []error
	for _, config := range *octyAlgoConfigs {
		switch config.AlgorithmName {
		case "rec":
			errs = api.rest.SetRecAlgorithmConfigurations(&config, credentials)
		case "churn":
			errs = api.rest.SetChurnAlgorithmConfigurations(&config, credentials)
		}
	}
	return errs
}

func (api Application) GetAlgorithmConfigs() (*[]c.OctyAlgorithmConfiguration, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetAlgorithmConfigurations(credentials)
}
