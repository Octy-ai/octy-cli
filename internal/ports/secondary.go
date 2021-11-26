package ports

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	t "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
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

	// templates
	CreateTemplates(templates *[]t.OctyMessageTemplate, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)
	GetTemplates(identifiers []string, credentials string) (*[]t.OctyMessageTemplate, []error)
	UpdateTemplates(templates *[]t.OctyMessageTemplate, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)
	DeleteTemplates(identifiers []string, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)

	// data upload
	GetResourceFormats(resourceType string) (map[string]string, map[int]string, error)
	UploadProfiles(profiles string, objectRowIDXMap *map[string]int, credentials string, prog *d.UploadProgess, progressChan chan<- d.UploadProgess)
}

type CredentialStorePort interface {
	SetOctyCredentials(pk string, sk string) error
	GetOctyCredentials() (string, error)
}
