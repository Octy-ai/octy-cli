package cli

import (
	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	e "github.com/Octy-ai/octy-cli/internal/application/domain/event_types"
	t "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
	s "github.com/Octy-ai/octy-cli/internal/application/domain/segments"
)

// applycmd models

type AccountConf struct {
	Kind           string                      `yaml:"kind"`
	Configurations c.OctyAccountConfigurations `yaml:"configurations"`
}

type AlgoConf struct {
	Kind           string                         `yaml:"kind"`
	Configurations []c.OctyAlgorithmConfiguration `yaml:"configurations"`
}

type EventTypes struct {
	Kind       string            `yaml:"kind"`
	EventTypes []e.OctyEventType `yaml:"eventTypeDefinitions"`
}

type Segments struct {
	Kind     string          `yaml:"kind"`
	Segments []s.OctySegment `yaml:"segmentDefinitions"`
}

type Templates struct {
	Kind      string                  `yaml:"kind"`
	Templates []t.OctyMessageTemplate `yaml:"templateDefinitions"`
}
