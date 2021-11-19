package messaging

type OctyMessageTemplate struct {
	TemplateID    string                 `json:"template_id,omitempty" yaml:"templateID,omitempty"`
	FriendlyName  string                 `json:"friendly_name" yaml:"friendlyName"`
	TemplateType  string                 `json:"template_type" yaml:"templateType"`
	Title         string                 `json:"title" yaml:"title"`
	Content       string                 `json:"content" yaml:"content"`
	RequiredData  []string               `json:"required_data" yaml:"requiredData"`
	DefaultValues map[string]string      `json:"default_values" yaml:"defaultValues"`
	Metadata      map[string]interface{} `json:"metadata,omitempty" yaml:"metadata"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	UpdatedAt     string                 `json:"updated_at,omitempty"`
	Status        string                 `json:"status,omitempty" yaml:"status"`
	ErrorMessage  string                 `json:"error_message,omitempty"`
}

//NewTemps : returns a pointer to a new slice of OctyMessageTemplate instances
func NewTemps() *[]OctyMessageTemplate {
	return &[]OctyMessageTemplate{}
}

//NewTemp : returns a pointer to a OctyMessageTemplate instance
func NewTemp() *OctyMessageTemplate {
	return &OctyMessageTemplate{}
}
