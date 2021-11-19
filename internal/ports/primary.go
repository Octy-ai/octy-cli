package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	s "github.com/Octy-ai/octy-cli/internal/application/domain/segments"
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

	// segments
	CreateSegments(segments *[]s.OctySegment) (*[]s.OctySegment, *[]s.OctySegment, []error)
	GetSegments(identifiers []string) (*[]s.OctySegment, []error)
	DeleteSegments(identifiers []string) (*[]s.OctySegment, *[]s.OctySegment, []error)
}
