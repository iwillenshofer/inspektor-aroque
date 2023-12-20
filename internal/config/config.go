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
func NewConfig() *Config {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("inspektor")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}
	return &config
}
