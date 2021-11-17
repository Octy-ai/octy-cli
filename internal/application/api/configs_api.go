package api

import c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"

//
// Public methods
//

func (api Application) SetAccountConfigs(octyAccConfig *c.OctyAccountConfigurations) error {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return err
	}
	return api.rest.SetAccountConfigurations(octyAccConfig, credentials)
}

func (api Application) GetAccountConfigs() (*c.OctyAccountConfigurations, error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, err
	}
	return api.rest.GetAccountConfigurations(credentials)
}

func (api Application) SetAlgorithmConfigs(octyAlgoConfigs *[]c.OctyAlgorithmConfiguration) error {

	var err error = nil

	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return err
	}

	for _, config := range *octyAlgoConfigs {
		switch config.AlgorithmName {
		case "rec":
			err = api.rest.SetRecAlgorithmConfigurations(&config, credentials)
		case "churn":
			err = api.rest.SetChurnAlgorithmConfigurations(&config, credentials)
		}
	}
	return err
}

func (api Application) GetAlgorithmConfigs() (*[]c.OctyAlgorithmConfiguration, error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, err
	}
	return api.rest.GetAlgorithmConfigurations(credentials)
}
