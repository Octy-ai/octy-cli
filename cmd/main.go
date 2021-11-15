package main

import (
	"log"

	"github.com/Octy-ai/octy-cli/internal/adapters/primary/cli"
	cs "github.com/Octy-ai/octy-cli/internal/adapters/secondary/credential_store"
	rest "github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest"
	"github.com/Octy-ai/octy-cli/internal/application/api"
)

func main() {

	restDrivenAdapter, err := rest.NewAdapter()
	if err != nil {
		log.Fatalf("failed to initialize rest driven adapter: %v", err)
	}
	csDrivenAdapter, err := cs.NewAdapter()
	if err != nil {
		log.Fatalf("failed to initialize crednetial store driven adapter: %v", err)
	}

	applicationAPI := api.NewApplication(restDrivenAdapter, csDrivenAdapter)

	cliAdapter := cli.NewAdapter(applicationAPI)

	cliAdapter.RegisterCommands()

	cliAdapter.ExecuteCMD()

}
