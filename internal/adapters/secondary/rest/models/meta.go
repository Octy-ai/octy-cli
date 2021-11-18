package models

type RequestMeta struct {
	RequestStatus string `json:"request_status"`
	Message       string `json:"message"`
	Count         int    `json:"count,omitempty"`
	Total         int    `json:"total,omitempty"`
}
