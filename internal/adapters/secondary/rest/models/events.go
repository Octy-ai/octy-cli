package models

import "encoding/json"

type OctyBatchCreateEventsResp struct {
	RequestMeta    RequestMeta           `json:"request_meta"`
	CreatedEvents  []CreatedEvent        `json:"created_events"`
	FailedToCreate []FailedToCreateEvent `json:"failed_to_create"`
}

type FailedToCreateEvent struct {
	EventType       string                 `json:"event_type"`
	EventProperties map[string]interface{} `json:"event_properties"`
	ProfileID       string                 `json:"profile_id"`
	ErrorMessage    string                 `json:"error_message"`
}

type CreatedEvent struct {
	EventID         string                 `json:"event_id"`
	ProfileID       string                 `json:"profile_id"`
	EventTypeID     string                 `json:"event_type_id"`
	EventType       string                 `json:"event_type"`
	EventProperties map[string]interface{} `json:"event_properties"`
	CreatedAt       string                 `json:"created_at"`
}

func UnmarshalOctyBatchCreateEventsResp(data []byte) (OctyBatchCreateEventsResp, error) {
	var r OctyBatchCreateEventsResp
	err := json.Unmarshal(data, &r)
	return r, err
}
