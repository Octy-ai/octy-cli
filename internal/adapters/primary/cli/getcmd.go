package cli

import "github.com/spf13/cobra"

type get struct {
	cmd *cobra.Command
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
			default:
				quit("error: you must specify a valid type of resource or configuration to get.", 1, nil)
			}
			return nil
		},
	}
	return g
}
