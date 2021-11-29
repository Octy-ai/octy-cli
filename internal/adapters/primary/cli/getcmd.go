package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

type get struct {
	cmd     *cobra.Command
	outpath string
}

func NewGetCmd(clia Adapter) *get {
	g := &get{}
	g.cmd = &cobra.Command{
		Use:   "get <type> <identifier> ...", // TODO: Add flag 'full' ot 'IDS' to either print full objects or just identifiers
		Args:  cobra.RangeArgs(1, 101),
		Short: "Get Octy configurations or resources.",
		Long:  `Get specififed Octy configurations or object definition resources.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			switch args[0] {
			case "accountconfig":
				getAccConfigController(clia)
			case "algorithmconfig":
				if len(args) < 2 {
					quit("error: accepts 2 arg(s), received 1. you must specify a valid type of configuration to get. 'rec' or 'churn'", 1, nil)
				}
				if args[1] != "rec" && args[1] != "churn" {
					quit("error: you must specify a valid type of configuration to get. 'rec' or 'churn'", 1, nil)
				}
				getAlgoConfigController(clia, args[1])
			case "eventtypes":
				getEventTypesController(clia, args[1:])
			case "segments":
				getSegmentsController(clia, args[1:])
			case "templates":
				getTemplatesController(clia, args[1:])
			case "churnreport":
				if g.outpath != "" {
					err := g.isValidDirectory(g.outpath)
					if err != nil {
						quit(err.Error(), 1, nil)
					}
				}
				getChurnPredictionReportController(clia, g.outpath)
			default:
				quit("error: you must specify a valid type of resource or configuration to get. Accepted: accountconfig, algorithmconfig, eventtypes, segments, templates, churnreport", 1, nil)
			}
			return nil
		},
	}
	g.registerFlags()
	return g
}

//
// Private methods
//

func (g *get) registerFlags() {
	g.cmd.Flags().StringVarP(&g.outpath, "outpath", "o", "", "Path to a directory where a markdown file containing a churn report will be stored (optional)")
}

// isValidDirectory: determines if the given directory exists and is valid
func (g *get) isValidDirectory(dir string) error {

	errStr := "please specify a valid directory to save the generated churn prediction report markdown file to. Example: ~/Desktop/"

	if g.outpath[len(g.outpath)-1:] != "/" {
		return errors.New(errStr)
	}
	_, err := os.Stat(g.outpath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return errors.New(errStr)
	}
	return nil
}
