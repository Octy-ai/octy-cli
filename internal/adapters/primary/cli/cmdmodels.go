package cli

import c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"

// applycmd models

type AccountConf struct {
	Kind           string                      `yaml:"kind"`
	Configurations c.OctyAccountConfigurations `yaml:"configurations"`
}

type AlgoConf struct {
	Kind           string                         `yaml:"kind"`
	Configurations []c.OctyAlgorithmConfiguration `yaml:"configurations"`
}
