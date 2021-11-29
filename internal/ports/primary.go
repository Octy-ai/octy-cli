package ports

import (
	cp "github.com/Octy-ai/octy-cli/internal/application/domain/churn_prediction_report"
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	t "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
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

	// templates
	CreateTemplates(templates *[]t.OctyMessageTemplate) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)
	GetTemplates(identifiers []string) (*[]t.OctyMessageTemplate, []error)
	UpdateTemplates(templates *[]t.OctyMessageTemplate) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)
	DeleteTemplates(identifiers []string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error)

	// data upload
	ValidateData(data *d.Data) (*[][]string, *map[string]int, []error)
	UploadData(resourceType string, objectIDXMap *map[string]int, content *[][]string, progressChan chan<- d.UploadProgess) []error

	// churn prediction report
	GetChurnReport() (*cp.OctyChurnPredictionReport, []error)
}
