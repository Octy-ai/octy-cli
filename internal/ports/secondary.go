package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
)

type RestPort interface {
	// auth
	Authenticate(pk string, sk string) error

	// config
	SetAccountConfigurations(octyAccConfig *c.OctyAccountConfigurations, credentials string) []error
	GetAccountConfigurations(credentials string) (*c.OctyAccountConfigurations, []error)
	SetRecAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) []error
	SetChurnAlgorithmConfigurations(octyAlgoConfig *c.OctyAlgorithmConfiguration, credentials string) []error
	GetAlgorithmConfigurations(credentials string) (*[]c.OctyAlgorithmConfiguration, []error)

	// event types
	CreateEventTypes(eventTypes *[]e.OctyEventType, credentials string) (*[]e.OctyEventType, *[]e.OctyEventType, []error)
	GetEventTypes(identifiers []string, credentials string) (*[]e.OctyEventType, []error)
	DeleteEventTypes(identifiers []string, credentials string) (*[]e.OctyEventType, *[]e.OctyEventType, []error)
}

type CredentialStorePort interface {
	SetOctyCredentials(pk string, sk string) error
	GetOctyCredentials() (string, error)
}
