package api

import (
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
)

//
// Public methods
//

func (api Application) CreateEventTypes(eventTypes *[]e.OctyEventType) (*[]e.OctyEventType, *[]e.OctyEventType, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.CreateEventTypes(eventTypes, credentials)
}

func (api Application) GetEventTypes(identifiers []string) (*[]e.OctyEventType, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetEventTypes(identifiers, credentials)
}

func (api Application) DeleteEventTypes(identifiers []string) (*[]e.OctyEventType, *[]e.OctyEventType, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.DeleteEventTypes(identifiers, credentials)
}
