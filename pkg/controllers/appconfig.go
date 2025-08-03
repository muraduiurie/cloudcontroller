package controllers

import (
	"fmt"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfigs struct {
	Controller ControllerConfigs `yaml:"controller"`
}

type CloudProviderConfigs struct {
	GCP GCPConfigs `yaml:"gcp"`
}

type GCPConfigs struct {
	GcpSaFilePath string `yaml:"gcpSaFilePath"`
}

type ControllerConfigs struct {
	CloudProviderConfigs `yaml:"cloudproviders"`
}

func LoadAppConfigs(log logr.Logger) (*AppConfigs, error) {
	log.Info("Loading app configs")
	var appConfigs AppConfigs
	// Read YAML file
	data, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	// Unmarshal YAML into struct
	if err = yaml.Unmarshal(data, &appConfigs); err != nil {
		return nil, fmt.Errorf("failed unmarshalling YAML: %w", err)
	}
	return &appConfigs, nil
}
