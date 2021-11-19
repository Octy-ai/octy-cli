package cli

import "github.com/spf13/cobra"

type delete struct {
	cmd *cobra.Command
}

func NewDeleteCmd(clia Adapter) *delete {
	d := &delete{}
	d.cmd = &cobra.Command{
		Use:   "delete <type> <identifier> ...",
		Args:  cobra.RangeArgs(1, 101),
		Short: "Delete Octy resources.",
		Long:  `Delete specififed Octy object definition resources.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			switch args[0] {
			case "eventtypes":
				if len(args) < 2 {
					quit("Error: accepts up to 100 event type identifiers, received 0. you must provide at least one event type identifier.", 1, nil)
				}
				deleteEventTypesController(clia, args[1:])
			case "segments":
				if len(args) < 2 {
					quit("Error: accepts up to 100 segment identifiers, received 0. you must provide at least one segment identifier.", 1, nil)
				}
				deleteSegmentsController(clia, args[1:])

			default:
				quit("you must specify a valid type of object definition resource to delete.", 1, nil)
			}
			return nil
		},
	}
	return d
}
