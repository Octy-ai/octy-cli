package versioning

type Version struct {
	ID          string  `json:"id"`
	VersionTag  string  `json:"version_tag"`
	VersionName string  `json:"version_name"`
	VersionInt  int64   `json:"version_int"`
	Assets      []Asset `json:"assets"`
	PublishedAt string  `json:"published_at"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

//NewVersion : returns a pointer to a Version instance
func NewVersion() *Version {
	return &Version{}
}
