package api

import (
	"github.com/Octy-ai/octy-cli/internal/ports"
)

type Application struct {
	rest   ports.RestPort
	cs     ports.CredentialStorePort
	upload Upload
}

func NewApplication(rest ports.RestPort, cs ports.CredentialStorePort, upload Upload) *Application {
	return &Application{
		rest:   rest,
		cs:     cs,
		upload: upload,
	}
}
