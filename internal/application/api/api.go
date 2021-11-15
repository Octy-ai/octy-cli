package api

import (
	"github.com/Octy-ai/octy-cli/internal/ports"
)

type Application struct {
	rest ports.RestPort
	cs   ports.CredentialStorePort
}

func NewApplication(rest ports.RestPort, cs ports.CredentialStorePort) *Application {
	return &Application{
		rest: rest,
		cs:   cs,
	}
}
