package api

import (
	"github.com/Octy-ai/octy-cli/internal/ports"
)

type Application struct {
	rest ports.RestPort
}

func NewApplication(rest ports.RestPort) *Application {
	return &Application{
		rest: rest,
	}
}
