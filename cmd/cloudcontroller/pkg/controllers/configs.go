package controllers

type Configs struct {
	Controller ControllerConfigs `yaml:"controller"`
}

type ControllerConfigs struct {
	GcpSaFilePath string `yaml:"gcpSaFilePath"`
}
