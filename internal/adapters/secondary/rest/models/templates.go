package models

import "encoding/json"

// ** Octy REST Request Models **

// ---

type OctyCreateMessageTemplatesReq struct {
	Templates []CreateTemplate `json:"templates"`
}

type CreateTemplate struct {
	FriendlyName  string                 `json:"friendly_name"`
	TemplateType  string                 `json:"template_type"`
	Title         string                 `json:"title"`
	Content       string                 `json:"content"`
	RequiredData  []string               `json:"required_data"`
	DefaultValues map[string]string      `json:"default_values"`
	Metadata      map[string]interface{} `json:"metadata"`
}

func (r *OctyCreateMessageTemplatesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctyUpdateMessageTemplatesReq struct {
	Templates []UpdateTemplate `json:"templates"`
}

type UpdateTemplate struct {
	FriendlyName  string                 `json:"friendly_name"`
	TemplateType  string                 `json:"template_type"`
	Title         string                 `json:"title"`
	Content       string                 `json:"content"`
	RequiredData  []string               `json:"required_data"`
	DefaultValues map[string]string      `json:"default_values"`
	Metadata      map[string]interface{} `json:"metadata"`
	Status        string                 `json:"status"`
	TemplateID    string                 `json:"template_id"`
}

func (r *OctyUpdateMessageTemplatesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctyDeleteMessageTemplatesReq struct {
	TemplateIDS []string `json:"template_ids"`
}

func (r *OctyDeleteMessageTemplatesReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Octy REST Response Models **

// ---

type OctyCreateMessageTemplatesResp struct {
	RequestMeta    RequestMeta          `json:"request_meta"`
	Templates      []CreateTemplateResp `json:"templates"`
	FailedToCreate []FailedToCreate     `json:"failed_to_create"`
}

type CreateTemplateResp struct {
	TemplateID    string                 `json:"template_id"`
	FriendlyName  string                 `json:"friendly_name"`
	TemplateType  string                 `json:"template_type"`
	Title         string                 `json:"title"`
	Content       string                 `json:"content"`
	RequiredData  []string               `json:"required_data"`
	DefaultValues map[string]string      `json:"default_values"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type FailedToCreate struct {
	FriendlyName string `json:"friendly_name"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyCreateMessageTemplatesResp(data []byte) (OctyCreateMessageTemplatesResp, error) {
	var r OctyCreateMessageTemplatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyGetMessageTemplatesResp struct {
	RequestMeta RequestMeta   `json:"request_meta"`
	Templates   []GetTemplate `json:"templates"`
}

type GetTemplate struct {
	FriendlyName  string                 `json:"friendly_name"`
	TemplateType  string                 `json:"template_type"`
	Title         string                 `json:"title"`
	Content       string                 `json:"content"`
	RequiredData  []string               `json:"required_data"`
	DefaultValues map[string]string      `json:"default_values"`
	Metadata      map[string]interface{} `json:"metadata"`
	Status        string                 `json:"status"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	TemplateID    string                 `json:"template_id"`
}

func UnmarshalOctyGetMessageTemplatesResp(data []byte) (OctyGetMessageTemplatesResp, error) {
	var r OctyGetMessageTemplatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyUpdateMessageTemplatesResp struct {
	RequestMeta    RequestMeta       `json:"request_meta"`
	Templates      []UpdatedTemplate `json:"templates"`
	FailedToUpdate []FailedToUpdate  `json:"failed_to_update"`
}

type UpdatedTemplate struct {
	TemplateID    string                 `json:"template_id"`
	FriendlyName  string                 `json:"friendly_name"`
	TemplateType  string                 `json:"template_type"`
	Title         string                 `json:"title"`
	Content       string                 `json:"content"`
	RequiredData  []string               `json:"required_data"`
	DefaultValues map[string]string      `json:"default_values"`
	Metadata      map[string]interface{} `json:"metadata"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

type FailedToUpdate struct {
	TemplateID   string `json:"template_id"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyUpdateMessageTemplatesResp(data []byte) (OctyUpdateMessageTemplatesResp, error) {
	var r OctyUpdateMessageTemplatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyDeleteMessageTemplatesResp struct {
	RequestMeta      RequestMeta       `json:"request_meta"`
	DeletedTemplates []DeletedTemplate `json:"deleted_templates"`
	FailedToDelete   []FailedToDelete  `json:"failed_to_delete"`
}

type DeletedTemplate struct {
	TemplateID string `json:"template_id"`
}

type FailedToDelete struct {
	TemplateID   string `json:"template_id"`
	ErrorMessage string `json:"error_message"`
}

func UnmarshalOctyDeleteMessageTemplatesResp(data []byte) (OctyDeleteMessageTemplatesResp, error) {
	var r OctyDeleteMessageTemplatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
