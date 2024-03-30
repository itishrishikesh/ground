package config

type masterConf struct {
	AppName  string
	Startup  startup
	Time     tym
	Database database
	Logging  loggingConfig
}

var config masterConf

func Store() *masterConf {
	return &config
}

func InitializeConfig() {
	// Should the startup variables be printed or not
	determineLaunchInfoPrintSetting()

	initializeViper()
	initializeStartupConfig()
	initializeTimeConfig()
	initializeDatabaseConfig()
	initializeLoggingConfig()
}
