package cli

import (
	"os"
	"strings"
	"time"

	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/output"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

type auth struct {
	cmd *cobra.Command

	publicKey string
	secretKey string
}

func NewAuthCmd(clia Adapter) *auth {
	a := &auth{}
	a.cmd = &cobra.Command{
		Use:   "auth",
		Short: "Configure your Octy API Authentication credentials.",
		Long: `Configure your Octy API Authentication credentials, granting this CLI limited access to your Octy account within the scope of the available options.
Your Octy credentials will be stored in this devices OS keychain. `,
		RunE: func(cmd *cobra.Command, args []string) error {
			if a.publicKey == "" {
				quit("Please specify --public-key flag e.g. '-p <octy public key>'.\nUse the -h flag for help using this command.", 1, nil)
			}
			if a.secretKey == "" {
				quit("Please specify --secret-key flag e.g. '-s <octy secret key>'.\nUse the -h flag for help using this command.", 1, nil)
			}
			spinner := output.StartNewSpinner("Authenticating account ...", os.Stdout)
			err := clia.api.SetOctyCredentials(a.publicKey, a.secretKey)
			if err != nil {
				if strings.Contains(err.Error(), "401") {
					quit("Could not authenticate with the provided credentials.", 1, spinner)
				}
				sentry.CaptureException(err)
				sentry.Flush(time.Second * 5)
				quit(globals.ErrUnknownError.Error(), 1, spinner)
			}
			quit("Set Octy credentials in OS keychain", 0, spinner)

			return nil
		},
	}

	a.registerFlags()

	return a
}

func (a *auth) registerFlags() {
	a.cmd.Flags().StringVarP(&a.publicKey, "public-key", "p", "", "Your Octy public key (required)")
	a.cmd.Flags().StringVarP(&a.secretKey, "secret-key", "s", "", "Your Octy secret key (required)")
}
