package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest/models"
	s "github.com/Octy-ai/octy-cli/internal/application/domain/segments"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/jinzhu/copier"
)

func (ha Adapter) CreateSegments(segments *[]s.OctySegment, credentials string) (*[]s.OctySegment, *[]s.OctySegment, []error) {

	var createdSegments = *s.NewSegs()
	var failedSegments = *s.NewSegs()

SegmentCreateLoop:
	for _, segment := range *segments {

		failedCreatedSeg := s.NewSeg()

		var es []models.EventSequenceEvent
		for _, e := range segment.EventSequence {
			es = append(es, models.EventSequenceEvent{
				EventType:       e.EventType,
				ExpTimeframe:    e.ExpTimeframe,
				ActionInaction:  e.ActionInaction,
				EventProperties: e.EventProperties,
			})
		}
		createSegmentReq := models.OctyCreateSegmentReq{
			SegmentName:          segment.SegmentName,
			SegmentType:          segment.SegmentType,
			SegmentSubType:       segment.SegmentSubType,
			SegmentTimeframe:     segment.SegmentTimeframe,
			EventSequence:        es,
			ProfilePropertyName:  segment.ProfilePropertyName,
			ProfilePropertyValue: segment.ProfilePropertyValue,
		}

		requestBody, err := createSegmentReq.Marshal()
		if err != nil {
			failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}

		req, err := http.NewRequest("POST", globals.CreateSegmentRoute, bytes.NewBuffer(requestBody))
		if err != nil {
			failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}

		req.Header.Add("Authorization", credentials)
		req.Header.Add("Content-Type", "application/json")

		resp, err := ha.httpClient.Do(req)
		if err != nil {
			failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}

		if resp.StatusCode > 201 {
			errResp, err := models.UnmarshalOctyErrorResp(body)
			if err != nil {
				failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
				failedCreatedSeg.SegmentName = segment.SegmentName
				failedSegments = append(failedSegments, *failedCreatedSeg)
				continue SegmentCreateLoop
			}
			failedCreatedSeg.ErrorMessage = errResp.Error.Errors[0].Message
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}

		createSegmentResp, err := models.UnmarshalOctyCreateSegmentResp(body)
		if err != nil {
			failedCreatedSeg.ErrorMessage = "Unknown error occurred when attmepting to create segment"
			failedCreatedSeg.SegmentName = segment.SegmentName
			failedSegments = append(failedSegments, *failedCreatedSeg)
			continue SegmentCreateLoop
		}
		failedCreatedSeg.SegmentID = createSegmentResp.SegmentID
		failedCreatedSeg.SegmentName = createSegmentResp.SegmentName
		createdSegments = append(createdSegments, *failedCreatedSeg)
	}
	return &createdSegments, &failedSegments, nil
}

func (ha Adapter) GetSegments(identifiers []string, credentials string) (*[]s.OctySegment, []error) {
	var cursor int = 0
	var url string
	foundSegments := *s.NewSegs()

	if len(identifiers) > 0 {
		var urlParams string
		for _, identifier := range identifiers {
			urlParams = urlParams + "," + identifier
		}
		url = globals.GetSegmentsRoute + "?ids=" + urlParams
	} else {
		url = globals.GetSegmentsRoute
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
			return nil, []error{errors.New("apierror[500]: unknown server error")}
		}

		getSegmentsResp, err := models.UnmarshalOctyGetSegmentsResp(body)
		if err != nil {
			return nil, []error{err}
		}

		for _, seg := range getSegmentsResp.Segments {
			foundSegment := s.NewSeg()
			copier.Copy(&foundSegment, &seg)
			foundSegments = append(foundSegments, *foundSegment)
		}

		if len(identifiers) > 0 {
			break PaginationLoop
		}
		rCursor, _ := strconv.Atoi(resp.Header.Get("cursor"))
		cursor += rCursor
		if getSegmentsResp.RequestMeta.Total >= cursor {
			break PaginationLoop
		}
	}

	return &foundSegments, nil
}

func (ha Adapter) DeleteSegments(identifiers []string, credentials string) (*[]s.OctySegment, *[]s.OctySegment, []error) {
	deleteSegmentsReq := models.OctyDeleteSegmentsReq{
		Segments: identifiers,
	}

	requestBody, err := deleteSegmentsReq.Marshal()
	if err != nil {
		return nil, nil, []error{err}
	}
	req, err := http.NewRequest("POST", globals.DeleteSegmentsRoute, bytes.NewBuffer(requestBody))
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
		return nil, nil, []error{errors.New("apierror[500]: unknown server error")}
	}

	deleteSegmentsResp, err := models.UnmarshalOctyDeleteSegmentsResp(body)
	if err != nil {
		return nil, nil, []error{err}
	}

	deletedSegments := s.NewSegs()
	failedDeletedSegments := s.NewSegs()

	copier.Copy(&deletedSegments, &deleteSegmentsResp.DeletedSegments)
	copier.Copy(&failedDeletedSegments, &deleteSegmentsResp.FailedToDelete)

	return deletedSegments, failedDeletedSegments, nil
}
