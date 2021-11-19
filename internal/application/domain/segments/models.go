package segments

type OctySegment struct {
	SegmentID            string               `json:"segment_id,omitempty"`
	SegmentName          string               `json:"segment_name,omitempty" yaml:"segmentName"`
	SegmentType          string               `json:"segment_type,omitempty" yaml:"segmentType"`
	SegmentSubType       int                  `json:"segment_sub_type,omitempty" yaml:"segmentSubtype"`
	SegmentTimeframe     int                  `json:"segment_timeframe,omitempty" yaml:"segmentTimeframe"`
	EventSequence        []EventSequenceEvent `json:"event_sequence,omitempty" yaml:"eventSequence"`
	ProfilePropertyName  string               `json:"profile_property_name,omitempty" yaml:"profilePropertyName,omitempty"`
	ProfilePropertyValue string               `json:"profile_property_value,omitempty" yaml:"profilePropertyValue,omitempty"`
	ProfileCount         int64                `json:"profile_count,omitempty"`
	CreatedAt            string               `json:"created_at,omitempty"`
	ErrorMessage         string               `json:"error_message,omitempty"`
}

type EventSequenceEvent struct {
	EventType       string                 `json:"event_type" yaml:"eventType"`
	ExpTimeframe    int                    `json:"exp_timeframe" yaml:"expTimeframe"`
	ActionInaction  string                 `json:"action_inaction" yaml:"actionInaction"`
	EventProperties map[string]interface{} `json:"event_properties" yaml:"eventProperties"`
}

//NewSegs : returns a pointer to a new slice of OctySegment instances
func NewSegs() *[]OctySegment {
	return &[]OctySegment{}
}

//NewSeg : returns a pointer to a OctySegment instance
func NewSeg() *OctySegment {
	return &OctySegment{}
}
