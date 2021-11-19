package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	s "github.com/Octy-ai/octy-cli/internal/application/domain/segments"
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

	// segments
	CreateSegments(segments *[]s.OctySegment, credentials string) (*[]s.OctySegment, *[]s.OctySegment, []error)
	GetSegments(identifiers []string, credentials string) (*[]s.OctySegment, []error)
	DeleteSegments(identifiers []string, credentials string) (*[]s.OctySegment, *[]s.OctySegment, []error)
}

type CredentialStorePort interface {
	SetOctyCredentials(pk string, sk string) error
	GetOctyCredentials() (string, error)
}
