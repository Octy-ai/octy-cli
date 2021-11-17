package models

import "encoding/json"

// ** Octy REST Request Models **

// ---

type OctySetAccConfigReq struct {
	ContactName         string `json:"contact_name"`
	ContactSurname      string `json:"contact_surname"`
	ContactEmailAddress string `json:"contact_email_address"`
	WebhookURL          string `json:"webhook_url"`
	AuthenticatedIDKey  string `json:"authenticated_id_key"`
}

func (r *OctySetAccConfigReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctySetRecAlgoConfigReq struct {
	AlgorithmName  string               `json:"algorithm_name"`
	Configurations RecReqConfigurations `json:"configurations"`
}

type RecReqConfigurations struct {
	RecommendInteractedItems bool     `json:"recommend_interacted_items"`
	ItemIDStopList           []string `json:"item_id_stop_list"`
	ProfileFeatures          []string `json:"profile_features"`
}

func (r *OctySetRecAlgoConfigReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type OctySetChurnAlgoConfigReq struct {
	AlgorithmName  string                 `json:"algorithm_name"`
	Configurations ChurnReqConfigurations `json:"configurations"`
}

type ChurnReqConfigurations struct {
	ProfileFeatures []string `json:"profile_features"`
}

func (r *OctySetChurnAlgoConfigReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Octy REST Response Models **

// ---

type OctyGetSetAccConfigResp struct {
	RequestMeta RequestMeta `json:"request_meta"`
	AccountData AccountData `json:"account_data"`
}

type AccountData struct {
	ContactName         string `json:"contact_name"`
	ContactSurname      string `json:"contact_surname"`
	ContactEmailAddress string `json:"contact_email_address"`
	WebhookURL          string `json:"webhook_url"`
	AuthenticatedIDKey  string `json:"authenticated_id_key"`
}

func UnmarshalOctyGetAccConfigResp(data []byte) (OctyGetSetAccConfigResp, error) {
	var r OctyGetSetAccConfigResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctySetAlgoConfigResp struct {
	RequestMeta    RequestMeta     `json:"request_meta"`
	Configurations []Configuration `json:"configurations"`
}

type Configuration struct {
	AlgorithmName  string         `json:"algorithm_name"`
	Configurations Configurations `json:"configurations"`
}

type Configurations struct {
	RecommendInteractedItems bool             `json:"recommend_interacted_items,omitempty"`
	ItemIDStopList           []ItemIDStopList `json:"item_id_stop_list,omitempty"`
	ProfileFeatures          []string         `json:"profile_features,omitempty"`
}

type ItemIDStopList struct {
	ItemID string `json:"item_id"`
}

func UnmarshalOctyGetAlgoConfigResp(data []byte) (OctySetAlgoConfigResp, error) {
	var r OctySetAlgoConfigResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
