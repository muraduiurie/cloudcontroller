package configmap

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigMap struct {
	Controller ControllerConfigs `yaml:"controller"`
	Common     CommonConfigs     `yaml:"common"`
}

type CommonConfigs struct {
	SecretsPath string `yaml:"secretsPath"`
}

type ControllerConfigs struct {
	GcpSaFilePath string `yaml:"gcpSaFilePath"`
}

func LoadConfigMap() (*ConfigMap, error) {
	var configMap ConfigMap
	// Read YAML file
	data, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	// Unmarshal YAML into struct
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return nil, fmt.Errorf("failed unmarshalling YAML: %w", err)
	}
	return &configMap, nil
}
