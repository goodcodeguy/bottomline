package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

// Configuration defines the structure for configuration
type Configuration struct {
	LogLevel string

	// Database Configuration
	DBHost    string
	DBPort    string
	DBName    string
	DBUser    string
	DBPass    string
	DBSSLMode string
}

// GetConfiguration from the environment
func GetConfiguration() Configuration {

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Printf("Error reading configuration file: %s", err.Error())
		os.Exit(500)
	}

	return configuration
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
