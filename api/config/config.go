// I decided to build this package up myself to provide a better dev experience,
// another solid alternative could be using the `GoDotEnv` package
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config is a struct used to capture configuration values from the yml file provided
type Config struct {
	DB     DbConfig     `yaml:"db"`
	AWS    AwsConfig    `yaml:"aws"`
	Router RouterConfig `yaml:"router"`
}

func getConfigFromDir(env string) ([]byte, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	yamlFilePath := filepath.Join(dir, fmt.Sprintf("config/%s.yml", env))

	yamlFile, err := ioutil.ReadFile(yamlFilePath)

	if err != nil {
		return nil, err
	}

	return yamlFile, nil
}

func getConfigFromExecutable(env string) ([]byte, error) {
	ex, err := os.Executable()

	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(ex)
	yamlFilePath := filepath.Join(dir, fmt.Sprintf("config/%s.yml", env))

	yamlFile, err := ioutil.ReadFile(yamlFilePath)

	if err != nil {
		return nil, err
	}

	return yamlFile, nil
}

// LoadDbConfig loads the DbConfig struct with values from the yaml
func New() Config {
	conf := Config{}
	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
	}

	// Attempt to load the config file using current directory
	// This will fail when running the code from an executable file
	yamlFile, err := getConfigFromDir(env)

	if yamlFile == nil {
		// Handles loading configuration when code is running from executable
		yamlFile, err = getConfigFromExecutable(env)
	}

	if err != nil {
		log.Fatalf("%s.yml not provided", env)
		log.Panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Fatalf("Config yaml unmarshalling failed: %v", err)
	}

	return conf
}
