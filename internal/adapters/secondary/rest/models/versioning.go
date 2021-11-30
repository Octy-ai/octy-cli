package models

import "encoding/json"

type OctyGetVersionInfoResp struct {
	RequestMeta     RequestMeta    `json:"request_meta"`
	ApplicationType string         `json:"application_type"`
	CurrentVersion  CurrentVersion `json:"current_version"`
}

type CurrentVersion struct {
	ID          string  `json:"id"`
	ReleaseID   int64   `json:"release_id"`
	VersionTag  string  `json:"version_tag"`
	VersionName string  `json:"version_name"`
	VersionInt  int64   `json:"version_int"`
	ChangeLog   string  `json:"change_log"`
	Assets      []Asset `json:"assets"`
	PublishedAt string  `json:"published_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Asset struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func UnmarshalOctyGetVersionInfoResp(data []byte) (OctyGetVersionInfoResp, error) {
	var r OctyGetVersionInfoResp
	err := json.Unmarshal(data, &r)
	return r, err
}
