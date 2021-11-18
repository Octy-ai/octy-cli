package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/Octy-ai/octy-cli/internal/ports"
	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/output"
	"github.com/briandowns/spinner"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("") },
}

//
// Public methods
//

func (clia Adapter) RegisterCommands() {
	rootCmd.AddCommand(NewAuthCmd(clia).cmd)
	rootCmd.AddCommand(NewApplyCmd(clia).cmd)
	rootCmd.AddCommand(NewGetCmd(clia).cmd)
	rootCmd.AddCommand(NewDeleteCmd(clia).cmd)
}

func (clia Adapter) ExecuteCMD() {
	rootCmd.Execute()
}

//
// Private functions
//

func init() {
	sentry.Init(sentry.ClientOptions{
		Dsn: globals.SentryDSN,
	})

	year, _, _ := time.Now().Date()
	fmt.Println("--")
	fmt.Printf("🐙 octy.ai © %v. \ncli-version: "+globals.CliVersion+" \napi-version: "+globals.ApiVersion+" \n", year)
	fmt.Println("--")
}

func quit(message string, code int, spinner *spinner.Spinner) {
	if spinner != nil {
		output.StopSpinner(spinner, "\n", code, os.Stdout)
	}
	if code == 1 {
		output.FPrint(message)
	} else {
		output.SPrint(message)
	}
	os.Exit(code)
}
