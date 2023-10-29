package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// New initializes the configuration based on the given environment.
//
// It loads the configuration file for the specified environment and sets the appropriate environment variables.
// The function also sets the default environment to "default" if none is provided.
func New(env string) {
	log.Info().Msg("loading config")

	viper.AutomaticEnv()
	// Replace env key
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	pwd, err := os.Getwd()
	if err != nil {
		log.Panic().Msg(fmt.Sprintf("get current directory failed: %v", err))
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(pwd)

	envFile := getEnvFile(env)
	viper.SetConfigFile(envFile)
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		pe := &fs.PathError{Op: "open", Path: envFile, Err: syscall.ENOENT}
		if ok := errors.As(err, &pe); !ok {
			log.Panic().Msg(fmt.Sprintf("read in config failed: %v", err))
		}
	}

	if env == "" {
		env = "default"
	}

	log.Info().
		Str("env", env).
		Str("goarch", runtime.GOARCH).
		Str("goos", runtime.GOOS).
		Str("version", runtime.Version()).
		Msg("load config successfully")
}

// getEnvFile returns the name of the environment file based on the provided environment.
//
// The env parameter is the name of the environment.
// It returns a string representing the name of the environment file.
func getEnvFile(env string) string {
	if env != "" {
		envFile := fmt.Sprintf(".env.%s", env)
		if checkEnvFileExist(envFile) {
			return envFile
		}
	}

	return ".env"
}

// checkEnvFileExist checks if the given environment file exists.
//
// Parameters:
// - envFile: a string representing the path to the environment file.
//
// Returns:
// - a boolean value indicating whether the environment file exists or not.
func checkEnvFileExist(envFile string) bool {
	_, err := os.Stat(envFile)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	return true
}
