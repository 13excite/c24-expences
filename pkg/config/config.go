// Package config contains the configuration for the service.
package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the main config of the service
type Config struct {
	InputDir   string           `yaml:"input_dir"`
	RunEvery   int              `yaml:"run_every"` // in hours
	Clickhouse ClickhouseConfig `yaml:"clickhouse"`
}

// ClickhouseConfig contains the configuration for the Clickhouse database
type ClickhouseConfig struct {
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"` // in seconds
}

// Defaults sets the default values for the configuration
func (conf *Config) Defaults() {
	conf.InputDir = "./input"
	conf.RunEvery = 24
	conf.Clickhouse = ClickhouseConfig{
		Address:  "localhost:9000",
		Database: "default",
		Username: "default",
		Password: "",
		Timeout:  5,
	}
}

// ReadConfigFile reading and parsing configuration yaml file
func (conf *Config) ReadConfigFile(configPath string) {
	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		log.Fatal(fmt.Errorf("could not unmarshal config %v", conf), err)
	}
}
