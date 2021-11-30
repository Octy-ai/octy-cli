package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	y "gopkg.in/yaml.v2"
	yk8 "sigs.k8s.io/yaml"
)

type apply struct {
	cmd *cobra.Command

	filePath string
	fileData []byte
}

func NewApplyCmd(clia Adapter) *apply {
	a := &apply{}
	a.cmd = &cobra.Command{
		Use:   "apply",
		Short: "Update configurations or Create/update Octy object definition resources.",
		Long: `Update configurations or Create/update Octy object definition resources. 
Configurations include account configurations and Octy algorithm configurations. Go here for more on configurations : https://octy.ai/docs/getting_started
Octy object definition resources are a set of structured properties that represent entities within the Octy ecosystem. 
For example: Segment definitions, Event type definitions and Message Template definitions.
To view example configuration and resource definitions yaml files please follow this link : https://github.com/Octy-ai/octy-cli/tree/master/examples/apply_cmd_yamls`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if a.filePath == "" {
				quit("Please specify --filepath flag e.g. '-f path/to/file.yaml'.", 1, nil)
			}

			err := a.getYamlFileData()
			if err != nil {
				quit(err.Error(), 1, nil)
			}

			kind, err := a.getKind()
			if err != nil {
				quit(err.Error(), 1, nil)
			}

			switch kind {
			case "accountConfigurations":
				var accountConf AccountConf
				if err := y.Unmarshal(a.fileData, &accountConf); err != nil {
					quit(err.Error(), 1, nil)
				}
				if err = accountConf.Validate(); err != nil {
					quit(err.Error(), 1, nil)
				}
				setAccConfigController(clia, accountConf)

			case "algorithmConfigurations":
				var algoConf AlgoConf
				if err := y.Unmarshal(a.fileData, &algoConf); err != nil {
					quit(err.Error(), 1, nil)
				}
				if err = algoConf.Validate(); err != nil {
					quit(err.Error(), 1, nil)
				}
				setAlgoConfigController(clia, algoConf)

			case "eventTypes":
				var eventTypes EventTypes
				if err := y.Unmarshal(a.fileData, &eventTypes); err != nil {
					quit(err.Error(), 1, nil)
				}
				if err = eventTypes.Validate(); err != nil {
					quit(err.Error(), 1, nil)
				}
				createEventTypesController(clia, eventTypes)
			case "segments":
				var segments Segments
				if err := y.Unmarshal(a.fileData, &segments); err != nil {
					quit(err.Error(), 1, nil)
				}
				if err = segments.Validate(); err != nil {
					quit(err.Error(), 1, nil)
				}
				createSegmentsController(clia, segments)
			case "templates":
				var templates Templates
				if err := y.Unmarshal(a.fileData, &templates); err != nil {
					quit(err.Error(), 1, nil)
				}
				if err = templates.Validate(); err != nil {
					quit(err.Error(), 1, nil)
				}
				createUpdateTemplatesController(clia, templates)

			default:
				quit("no valid resource types found in specified yaml file. Accepted: accountConfigurations, algorithmConfigurations, eventTypes, segments, templates.\nUse the -h flag for help using this command.", 1, nil)
			}

			return nil
		},
	}

	a.registerFlags()

	return a
}

//
// Private methods
//

func (a *apply) registerFlags() {
	a.cmd.Flags().StringVarP(&a.filePath, "filepath", "f", "", "Path to the YAML file that contains configurations or Octy object definition resources (required)")
}

func (a *apply) getYamlFileData() error {
	f, err := ioutil.ReadFile(a.filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %s please ensure that file file exists and is readable", err)
	}
	a.fileData = f
	return nil
}

func (a *apply) getKind() (string, error) {
	yamlJson, err := yk8.YAMLToJSON(a.fileData)
	if err != nil {
		return "nil", fmt.Errorf("error reading YAML file: %s", err)
	}
	yamlMap := make(map[string]interface{})
	err = json.Unmarshal(yamlJson, &yamlMap)
	if err != nil {
		return "", fmt.Errorf("error reading YAML file: %s", err)
	}
	if !keyExists(yamlMap, "kind") {
		return "", errors.New("invalid configuration yaml file provided. Missing required key: 'kind'")
	}
	return yamlMap["kind"].(string), nil
}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}
