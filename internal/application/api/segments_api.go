package api

import (
	s "github.com/Octy-ai/octy-cli/internal/application/domain/segments"
)

//
// Public methods
//

func (api Application) CreateSegments(segments *[]s.OctySegment) (*[]s.OctySegment, *[]s.OctySegment, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.CreateSegments(segments, credentials)
}

func (api Application) GetSegments(identifiers []string) (*[]s.OctySegment, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, []error{err}
	}
	return api.rest.GetSegments(identifiers, credentials)
}

func (api Application) DeleteSegments(identifiers []string) (*[]s.OctySegment, *[]s.OctySegment, []error) {
	credentials, err := api.cs.GetOctyCredentials()
	if err != nil {
		return nil, nil, []error{err}
	}
	return api.rest.DeleteSegments(identifiers, credentials)
}
