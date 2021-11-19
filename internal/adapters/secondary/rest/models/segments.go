package models

import "encoding/json"

// ** Octy REST Request Models **

// ---

type OctyCreateSegmentReq struct {
	SegmentName          string               `json:"segment_name"`
	SegmentType          string               `json:"segment_type"`
	SegmentSubType       int                  `json:"segment_sub_type"`
	SegmentTimeframe     int                  `json:"segment_timeframe"`
	EventSequence        []EventSequenceEvent `json:"event_sequence"`
	ProfilePropertyName  string               `json:"profile_property_name,omitempty"`
	ProfilePropertyValue string               `json:"profile_property_value,omitempty"`
}

type EventSequenceEvent struct {
	EventType       string                 `json:"event_type"`
	ExpTimeframe    int                    `json:"exp_timeframe"`
	ActionInaction  string                 `json:"action_inaction"`
	EventProperties map[string]interface{} `json:"event_properties"`
}

func (r *OctyCreateSegmentReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctyDeleteSegmentsReq struct {
	Segments []string `json:"segments"`
}

func (r *OctyDeleteSegmentsReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Octy REST Response Models **

// ---

type OctyCreateSegmentResp struct {
	RequestMeta    RequestMeta `json:"request_meta"`
	SegmentID      string      `json:"segment_id"`
	SegmentName    string      `json:"segment_name"`
	SegmentType    string      `json:"segment_type"`
	SegmentSubType int64       `json:"segment_sub_type"`
	SegmentStatus  string      `json:"segment_status"`
}

func UnmarshalOctyCreateSegmentResp(data []byte) (OctyCreateSegmentResp, error) {
	var r OctyCreateSegmentResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyDeleteSegmentsResp struct {
	RequestMeta     RequestMeta            `json:"request_meta"`
	DeletedSegments []DeletedSegment       `json:"deleted_segments"`
	FailedToDelete  []FailedDeletedSegment `json:"failed_to_delete"`
}

type DeletedSegment struct {
	SegmentID string `json:"segment_id"`
}

type FailedDeletedSegment struct {
	SegmentID    string `json:"segment_id"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyDeleteSegmentsResp(data []byte) (OctyDeleteSegmentsResp, error) {
	var r OctyDeleteSegmentsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyGetSegmentsResp struct {
	RequestMeta RequestMeta `json:"request_meta"`
	Segments    []Segment   `json:"segments"`
}

type Segment struct {
	SegmentID            string          `json:"segment_id"`
	SegmentName          string          `json:"segment_name"`
	SegmentType          string          `json:"segment_type"`
	SegmentSubType       int64           `json:"segment_sub_type"`
	SegmentTimeframe     int64           `json:"segment_timeframe"`
	EventSequence        []EventSequence `json:"event_sequence"`
	ProfilePropertyName  string          `json:"profile_property_name"`
	ProfilePropertyValue string          `json:"profile_property_value"`
	ProfileCount         int64           `json:"profile_count"`
	CreatedAt            string          `json:"created_at"`
}

type EventSequence struct {
	EventType       string                 `json:"event_type"`
	ExpTimeframe    int64                  `json:"exp_timeframe"`
	ActionInaction  string                 `json:"action_inaction"`
	EventProperties map[string]interface{} `json:"event_properties"`
}

func UnmarshalOctyGetSegmentsResp(data []byte) (OctyGetSegmentsResp, error) {
	var r OctyGetSegmentsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
