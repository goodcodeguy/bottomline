package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// ConfigurationSpec defines the structure for configuration
type ConfigurationSpec struct {
	LogLevel string

	// Database Configuration
	DBHost    string
	DBPort    string
	DBName    string
	DBUser    string
	DBPass    string
	DBSSLMode string
}

// Configuration from the environment
func Configuration() ConfigurationSpec {
	var configuration ConfigurationSpec

	err := envconfig.Process("bottomline", &configuration)
	if err != nil {
		log.Fatalf("Error Initializing Configuration: %s", err.Error())
	}

	return configuration
}
