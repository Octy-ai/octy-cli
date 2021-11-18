package models

import "encoding/json"

// ** Octy REST Request Models **

// ---

type OctyCreateEventTypesReq struct {
	EventTypes []EventType `json:"event_types"`
}

type EventType struct {
	EventType       string   `json:"event_type"`
	EventProperties []string `json:"event_properties"`
}

func (r *OctyCreateEventTypesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctyDeleteEventTypesReq struct {
	EventTypeIDS []string `json:"event_type_ids"`
}

func (r *OctyDeleteEventTypesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Octy REST Response Models **

// ---

type OctyCreateEventTypesResp struct {
	RequestMeta    RequestMeta        `json:"request_meta"`
	EventTypes     []CreatedEventType `json:"event_types"`
	FailedToCreate []FailedEventType  `json:"failed_to_create"`
}

type CreatedEventType struct {
	EventTypeID     string   `json:"event_type_id"`
	EventType       string   `json:"event_type"`
	EventProperties []string `json:"event_properties"`
}

type FailedEventType struct {
	EventType    string `json:"event_type"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyCreateEventTypesResp(data []byte) (OctyCreateEventTypesResp, error) {
	var r OctyCreateEventTypesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyGetEventTypesResp struct {
	RequestMeta RequestMeta    `json:"request_meta"`
	EventTypes  []GetEventType `json:"event_types"`
}

type GetEventType struct {
	EventTypeID     string   `json:"event_type_id"`
	EventType       string   `json:"event_type"`
	EventProperties []string `json:"event_properties"`
	CreatedAt       string   `json:"created_at"`
}

func UnmarshalOctyGetEventTypesResp(data []byte) (OctyGetEventTypesResp, error) {
	var r OctyGetEventTypesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyDeleteEventTypesResp struct {
	RequestMeta       RequestMeta             `json:"request_meta"`
	DeletedEventTypes []DeletedEventType      `json:"deleted_event_types"`
	FailedToDelete    []FailedDeleteEventType `json:"failed_to_delete"`
}

type DeletedEventType struct {
	EventTypeID string `json:"event_type_id"`
}

type FailedDeleteEventType struct {
	EventTypeID  string `json:"event_type_id"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyDeleteEventTypesResp(data []byte) (OctyDeleteEventTypesResp, error) {
	var r OctyDeleteEventTypesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
