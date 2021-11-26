package data_upload

import (
	"sync"

	"github.com/schollz/progressbar/v3"
)

// Upload logic
type Upload struct {
}

//NewUpload: returns a pointer to a Upload instance
func NewUpload() *Upload {
	return &Upload{}
}

// ---

// Basic structure used to pass data to api.

type Data struct {
	ResourceType string
	Data         *[]byte
}

//NewData : returns a pointer to a Data instance
func NewData() *Data {
	return &Data{}
}

// ---

// Tracking upload progress across goroutines

type Failed struct {
	ErrorMessage string
	RowIDX       int
}

//NewFailed : returns a pointer to a Failed instance
func NewFailed() *Failed {
	return &Failed{}
}

type UploadProgess struct {
	Mutex       sync.RWMutex
	Bar         *progressbar.ProgressBar
	Complete    int64    // total number of successfully created objects
	Failed      []Failed // list of failed to create objects with errors and index of failed objects
	Total       int64    // total number of objects to be processed
	TotalChunks int64    // total number of chunks to be uploaded
}

//UploadProgess : returns a pointer to a UploadProgess instance
func NewUploadProgess() *UploadProgess {
	return &UploadProgess{}
}
