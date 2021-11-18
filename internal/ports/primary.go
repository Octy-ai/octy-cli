package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
)

type APIPort interface {

	// auth
	SetOctyCredentials(pk string, sk string) error

	// configs
	SetAccountConfigs(octyAccConfig *c.OctyAccountConfigurations) []error
	GetAccountConfigs() (*c.OctyAccountConfigurations, []error)
	SetAlgorithmConfigs(octyAlgoConfigs *[]c.OctyAlgorithmConfiguration) []error
	GetAlgorithmConfigs() (*[]c.OctyAlgorithmConfiguration, []error)

	// event types
	CreateEventTypes(eventTypes *[]e.OctyEventType) (*[]e.OctyEventType, *[]e.OctyEventType, []error)
	GetEventTypes(identifiers []string) (*[]e.OctyEventType, []error)
	DeleteEventTypes(identifiers []string) (*[]e.OctyEventType, *[]e.OctyEventType, []error)
}
