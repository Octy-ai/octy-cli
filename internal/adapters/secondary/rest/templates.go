package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	t "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

func (ha Adapter) CreateTemplates(templates *[]t.OctyMessageTemplate, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {

	var ts []models.CreateTemplate
	for _, t := range *templates {
		ts = append(ts, models.CreateTemplate{
			FriendlyName:  t.FriendlyName,
			TemplateType:  t.TemplateType,
			Title:         t.Title,
			Content:       t.Content,
			DefaultValues: t.DefaultValues,
			Metadata:      t.Metadata,
		})
	}
	createTemplatesReq := models.OctyCreateMessageTemplatesReq{
		Templates: ts,
	}

	requestBody, err := createTemplatesReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.CreateTemplatesRoute, bytes.NewBuffer(requestBody))
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

	createTemplatesResp, err := models.UnmarshalOctyCreateMessageTemplatesResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	createdTemplates := t.NewTemps()
	failedTemplates := t.NewTemps()

	copier.Copy(&createdTemplates, &createTemplatesResp.Templates)
	copier.Copy(&failedTemplates, &createTemplatesResp.FailedToCreate)

	return createdTemplates, failedTemplates, nil
}

func (ha Adapter) GetTemplates(identifiers []string, credentials string) (*[]t.OctyMessageTemplate, []error) {

	var cursor int = 0
	var url string
	var foundTemplates = *t.NewTemps()

	if len(identifiers) > 0 {
		var urlParams string
		for _, identifier := range identifiers {
			urlParams = urlParams + "," + identifier
		}
		url = globals.GetTemplatesRoute + "?ids=" + urlParams
	} else {
		url = globals.GetTemplatesRoute
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

		getTemplatesResp, err := models.UnmarshalOctyGetMessageTemplatesResp(body)
		if err != nil {
			return nil, []error{err}
		}

		for _, te := range getTemplatesResp.Templates {
			nte := t.NewTemp()
			copier.Copy(&nte, &te)

			foundTemplates = append(foundTemplates, *nte)
		}

		if len(identifiers) > 0 {
			break PaginationLoop
		}
		rCursor, _ := strconv.Atoi(resp.Header.Get("cursor"))
		cursor += rCursor

		if getTemplatesResp.RequestMeta.Total >= cursor {
			break PaginationLoop
		}
	}

	return &foundTemplates, nil
}

func (ha Adapter) UpdateTemplates(templates *[]t.OctyMessageTemplate, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {

	var ts []models.UpdateTemplate
	for _, t := range *templates {
		ts = append(ts, models.UpdateTemplate{
			TemplateID:    t.TemplateID,
			FriendlyName:  t.FriendlyName,
			TemplateType:  t.TemplateType,
			Title:         t.Title,
			Content:       t.Content,
			DefaultValues: t.DefaultValues,
			Metadata:      t.Metadata,
			Status:        t.Status,
		})
	}
	updateTemplatesReq := models.OctyUpdateMessageTemplatesReq{
		Templates: ts,
	}

	requestBody, err := updateTemplatesReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.UpdateTemplatesRoute, bytes.NewBuffer(requestBody))
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

	updateTemplatesResp, err := models.UnmarshalOctyUpdateMessageTemplatesResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	updateTemplates := t.NewTemps()
	failedTemplates := t.NewTemps()

	copier.Copy(&updateTemplates, &updateTemplatesResp.Templates)
	copier.Copy(&failedTemplates, &updateTemplatesResp.FailedToUpdate)

	return updateTemplates, failedTemplates, nil
}

func (ha Adapter) DeleteTemplates(identifiers []string, credentials string) (*[]t.OctyMessageTemplate, *[]t.OctyMessageTemplate, []error) {

	deleteTemplatesReq := models.OctyDeleteMessageTemplatesReq{
		TemplateIDS: identifiers,
	}

	requestBody, err := deleteTemplatesReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.DeleteTemplatesRoute, bytes.NewBuffer(requestBody))
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

	deleteTemplatesResp, err := models.UnmarshalOctyDeleteMessageTemplatesResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	deleteTemplates := t.NewTemps()
	failedDeleteTemplates := t.NewTemps()

	copier.Copy(&deleteTemplates, &deleteTemplatesResp.DeletedTemplates)
	copier.Copy(&failedDeleteTemplates, &deleteTemplatesResp.FailedToDelete)

	return deleteTemplates, failedDeleteTemplates, nil
}
