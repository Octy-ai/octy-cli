package api

import (
	t "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
)

//
// Public methods
//

func (api Application) CreateTemplates(templates *[]t.OctyMessageTemplate) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.CreateTemplates(templates, credentials)
}

func (api Application) GetTemplates(identifiers []string) (*[]t.OctyMessageTemplate, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetTemplates(identifiers, credentials)
}

func (api Application) UpdateTemplates(templates *[]t.OctyMessageTemplate) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.UpdateTemplates(templates, credentials)
}

func (api Application) DeleteTemplates(identifiers []string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.DeleteTemplates(identifiers, credentials)
}
