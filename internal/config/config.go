package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config is a struct that holds the configuration for the application
type Config struct {
	// Server is a struct that holds the configuration for the server
	Server ServerConfig
}

// ServerConfig is a struct that holds the configuration for the server
type ServerConfig struct {
	// Host is the host that the server will listen on
	Host string
	// Port is the port that the server will listen on
	Port int
}

// NewConfig returns a new Config struct
func NewConfig() (Config, error) {

	// Set the default values for the configuration
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)

	// Set the configuration file name
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Configuration file not found")
		} else {
			return Config{}, err
		}
	}

	// Set the environment variables prefix
	viper.SetEnvPrefix("inspektor")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Unmarshal the configuration file into the Config struct
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
