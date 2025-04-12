package controllers

import (
	"os"
	"testing"
)

func TestLoadAppConfigs(t *testing.T) {
	os.Setenv("CONFIG_PATH", "./testdata/appconfig.yaml")
	appConfigs, err := LoadAppConfigs(testLogger())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if appConfigs.Controller.CloudProviderConfigs.GCP.GcpSaFilePath != "/testdata/gcp_credentials" {
		t.Fatalf("expected gcpSaFilePath to be testdata/gcp_credentials, got %s", appConfigs.Controller.CloudProviderConfigs.GCP.GcpSaFilePath)
	}
}
