package config

import (
	"github.com/spf13/viper"
)

// AppName returns the name of the application.
func AppName() string {
	return viper.GetString("APP_NAME")
}

// AppPort returns the value of the APP_PORT configuration variable as an integer.
func AppPort() int {
	return viper.GetInt("APP_PORT")
}

// PprofEnabled returns a boolean value indicating whether PPROF_ENABLED is true
func PprofEnabled() bool {
	return viper.GetBool("PPROF_ENABLED")
}

// PprofPort returns the value of the "PPROF_PORT" configuration setting.
func PprofPort() int {
	return viper.GetInt("PPROF_PORT")
}

// LogPayload returns a boolean value indicating whether to log the payload.
func LogPayload() bool {
	return viper.GetBool("LOG_PAYLOAD")
}
