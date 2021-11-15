package rest

import (
	"net/http"
)

type Adapter struct {
	httpClient *http.Client
}

func NewAdapter() (*Adapter, error) {
	httpClient := &http.Client{}
	return &Adapter{httpClient: httpClient}, nil
}
