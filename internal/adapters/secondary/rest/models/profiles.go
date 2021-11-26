package models

import "encoding/json"

// ** Octy REST Response Models **

// ---

type OctyCreateProfilesResp struct {
	RequestMeta    RequestMeta           `json:"request_meta"`
	Profiles       []Profile             `json:"profiles"`
	FailedToCreate []FailedCreateProfile `json:"failed_to_create"`
}

type FailedCreateProfile struct {
	CustomerID   string `json:"customer_id"`
	ErrorMessage string `json:"error_message"`
}

type Profile struct {
	ProfileID    string                 `json:"profile_id"`
	CustomerID   string                 `json:"customer_id"`
	ProfileData  map[string]interface{} `json:"profile_data"`
	PlatformInfo map[string]interface{} `json:"platform_info"`
	HasCharged   bool                   `json:"has_charged"`
	Status       string                 `json:"status"`
	CreatedAt    string                 `json:"created_at"`
}

func UnmarshalOctyCreateProfilesResp(data []byte) (OctyCreateProfilesResp, error) {
	var r OctyCreateProfilesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
