package cli

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	d "github.com/Octy-ai/octy-cli/internal/application/domain/data_upload"
	"github.com/spf13/cobra"
)

type upload struct {
	cmd *cobra.Command

	filePath string
	data     *d.Data
}

func NewUploadCmd(clia Adapter) *upload {
	u := &upload{}
	u.cmd = &cobra.Command{
		Use:   "upload <type>",
		Args:  cobra.ExactArgs(1),
		Short: "Upload Octy resource data.",
		Long: `Upload Octy resource data.
Octy resource data is used in analytical and machine-learing training jobs. Resources in the context of the retention toolchain include the following:
profiles -- At the centre of Octy's toolchain is the customer profile. A profile is an anonymised representation of each customer. 
items -- Structured representations of the products your business sells.
events -- Actions performed by profiles.
For more information about resources, please see the documentation : https://octy.ai/docs/creating_resources
To view example resource csv files please follow this link : https://github.com/Octy-ai/octy-cli/tree/master/examples/upload_cmd_csvs
Accepted types:
- profiles
- items
- events`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if u.filePath == "" {
				quit("Please specify --filepath flag e.g. '-f path/to/file.csv'", 1, nil)
			}

			switch args[0] {
			case "profiles", "items", "events":
				u.data = d.NewData()
				u.data.ResourceType = args[0]
				err := u.getCsvFileData()
				if err != nil {
					quit(err.Error(), 1, nil)
				}
				uploadDataController(clia, u.data)
			default:
				quit("invalid resource type specififed. Accepted: profiles, items, events.\nUse the -h flag for help using this command.", 1, nil)
			}

			return nil
		},
	}

	u.registerFlags()

	return u
}

//
// Private methods
//

func (u *upload) registerFlags() {
	u.cmd.Flags().StringVarP(&u.filePath, "filepath", "f", "", "Path to the CSV file that contains Octy resource data (required)")
}

func (u *upload) getCsvFileData() error {

	extension := filepath.Ext(u.filePath)
	if extension != ".csv" {
		return fmt.Errorf("invalid file extension: %s. Expected .csv", extension)
	}
	data, err := ioutil.ReadFile(u.filePath)
	if err != nil {
		return fmt.Errorf("error reading CSV file: %s", err)
	}
	u.data.Data = &data
	return nil
}
