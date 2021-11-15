package main

import (
	"log"

	"github.com/Octy-ai/octy-cli/internal/adapters/primary/cli"
	rest "github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest"
	"github.com/Octy-ai/octy-cli/internal/application/api"
)

func main() {

	restDrivenAdapter, err := rest.NewAdapter()
	if err != nil {
		log.Fatalf("failed to initialize rest driven adapter: %v", err)
	}

	applicationAPI := api.NewApplication(restDrivenAdapter)

	cliAdapter := cli.NewAdapter(applicationAPI)

	cliAdapter.RegisterCommands()

	cliAdapter.ExecuteCMD()

}
