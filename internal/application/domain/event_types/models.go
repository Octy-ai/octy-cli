package eventtypes

// octy event types domain models

type OctyEventType struct {
	EventTypeID     string   `json:"event_type_id,omitempty"`
	EventType       string   `json:"event_type" yaml:"eventType"`
	EventProperties []string `json:"event_properties" yaml:"eventProperties"`
	CreatedAt       string   `json:"created_at,omitempty"`
	ErrorMessage    string   `json:"error_message,omitempty"`
}

//NewETS : returns a pointer to a new slice of OctyEventType instances
func NewETS() *[]OctyEventType {
	return &[]OctyEventType{}
}

//NewETS : returns a pointer to a OctyEventType instance
func NewET() *OctyEventType {
	return &OctyEventType{}
}
