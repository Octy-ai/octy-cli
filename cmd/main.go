package main

import (
	"log"

	"github.com/Octy-ai/octy-cli/internal/adapters/primary/cli"
	cs "github.com/Octy-ai/octy-cli/internal/adapters/secondary/credential_store"
	rest "github.com/Octy-ai/octy-cli/internal/adapters/secondary/rest"
	"github.com/Octy-ai/octy-cli/internal/application/api"
	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
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

	upload := d.NewUpload()

	applicationAPI := api.NewApplication(restDrivenAdapter, csDrivenAdapter, upload)

	cliAdapter := cli.NewAdapter(applicationAPI)

	cliAdapter.RegisterCommands()

	cliAdapter.VersionAssesment()

	cliAdapter.ExecuteCMD()
}
