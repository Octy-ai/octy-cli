package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
)

type RestPort interface {
	// auth
	Authenticate(pk string, sk string) error

	// config
	SetAccountConfigurations(octyAccConfig *c.OctyAccountConfigurations, credentials string) error
	GetAccountConfigurations(credentials string) (*c.OctyAccountConfigurations, error)
	SetRecAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) error
	SetChurnAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) error
	GetAlgorithmConfigurations(credentials string) (*[]c.OctyAlgorithmConfiguration, error)
}

type CredentialStorePort interface {
	SetOctyCredentials(pk string, sk string) error
	GetOctyCredentials() (string, error)
}
