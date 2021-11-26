package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

func (ha Adapter) CreateEventTypes(eventTypes *[]e.OctyEventType, credentials string) (*[]e.OctyEventType, *[]e.OctyEventType, []error) {

	var types []models.EventType
	for _, e := range *eventTypes {
		types = append(types, models.EventType{
			EventType:       e.EventType,
			EventProperties: e.EventProperties,
		})
	}
	createEventTypesReq := models.OctyCreateEventTypesReq{
		EventTypes: types,
	}

	requestBody, err := createEventTypesReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.CreateEventTypesRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, nil, []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, []error{err}
	}

	switch {
	case resp.StatusCode > 201 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, nil, []error{err}
		}
		return nil, nil, models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return nil, nil, []error{errors.New("apierror[500]:: unknown server error")}
	}

	createEventTypesResp, err := models.UnmarshalOctyCreateEventTypesResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	createdEventTypes := e.NewETS()
	failedEventTypes := e.NewETS()

	copier.Copy(&createdEventTypes, &createEventTypesResp.EventTypes)
	copier.Copy(&failedEventTypes, &createEventTypesResp.FailedToCreate)

	return createdEventTypes, failedEventTypes, nil
}

func (ha Adapter) GetEventTypes(identifiers []string, credentials string) (*[]e.OctyEventType, []error) {

	var cursor int = 0
	var url string
	var foundEventTypes = *e.NewETS()

	if len(identifiers) > 0 {
		var urlParams string
		for _, identifier := range identifiers {
			urlParams = urlParams + "," + identifier
		}
		url = globals.GetEventTypesRoute + "?ids=" + urlParams
	} else {
		url = globals.GetEventTypesRoute
	}

PaginationLoop:
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, []error{err}
		}
		req.Header.Add("Authorization", credentials)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("cursor", strconv.Itoa(cursor))

		resp, err := ha.httpClient.Do(req)
		if err != nil {
			return nil, []error{err}
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, []error{err}
		}

		switch {
		case resp.StatusCode > 200 && resp.StatusCode < 500:
			errResp, err := models.UnmarshalOctyErrorResp(body)
			if err != nil {
				return nil, []error{err}
			}
			return nil, models.ParseErrors(errResp)
		case resp.StatusCode >= 500:
			return nil, []error{errors.New("apierror[500]:: unknown server error")}
		}

		getEventTypesResp, err := models.UnmarshalOctyGetEventTypesResp(body)
		if err != nil {
			return nil, []error{err}
		}

		for _, etr := range getEventTypesResp.EventTypes {
			et := e.NewET()
			copier.Copy(&et, &etr)

			foundEventTypes = append(foundEventTypes, *et)
		}

		if len(identifiers) > 0 {
			break PaginationLoop
		}
		rCursor, _ := strconv.Atoi(resp.Header.Get("cursor"))
		cursor += rCursor

		if getEventTypesResp.RequestMeta.Total >= cursor {
			break PaginationLoop
		}
	}

	return &foundEventTypes, nil
}

func (ha Adapter) DeleteEventTypes(identifiers []string, credentials string) (*[]e.OctyEventType, *[]e.OctyEventType, []error) {

	deleteEventTypesReq := models.OctyDeleteEventTypesReq{
		EventTypeIDS: identifiers,
	}

	requestBody, err := deleteEventTypesReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.DeleteEventTypesRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, []error{err}
	}

	req.Header.Add("Authorization", credentials)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, nil, []error{err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, []error{err}
	}

	switch {
	case resp.StatusCode > 200 && resp.StatusCode < 500:
		errResp, err := models.UnmarshalOctyErrorResp(body)
		if err != nil {
			return nil, nil, []error{err}
		}
		return nil, nil, models.ParseErrors(errResp)
	case resp.StatusCode >= 500:
		return nil, nil, []error{errors.New("apierror[500]:: unknown server error")}
	}

	deleteEventTypesResp, err := models.UnmarshalOctyDeleteEventTypesResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	deletedEventTypes := e.NewETS()
	failedDeletedEventTypes := e.NewETS()

	copier.Copy(&deletedEventTypes, &deleteEventTypesResp.DeletedEventTypes)
	copier.Copy(&failedDeletedEventTypes, &deleteEventTypesResp.FailedToDelete)

	return deletedEventTypes, failedDeletedEventTypes, nil
}
