package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
)

type APIPort interface {

	// auth
	SetOctyCredentials(pk string, sk string) error

	// configs
	SetAccountConfigs(octyAccConfig *c.OctyAccountConfigurations) error
	GetAccountConfigs() (*c.OctyAccountConfigurations, error)
	SetAlgorithmConfigs(octyAlgoConfigs *[]c.OctyAlgorithmConfiguration) error
	GetAlgorithmConfigs() (*[]c.OctyAlgorithmConfiguration, error)
}
