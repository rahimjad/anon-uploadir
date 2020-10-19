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
	DB  DbConfig  `yaml:"db"`
	AWS AwsConfig `yaml:"aws"`
}

// LoadDbConfig loads the DbConfig struct with values from the yaml
func New() Config {
	conf := Config{}

	pwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Could not get root path")
	}

	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
	}

	yamlFilePath := filepath.Join(pwd, fmt.Sprintf("config/%s.yml", env))
	yamlFile, err := ioutil.ReadFile(yamlFilePath)

	if err != nil {
		log.Fatalf("%s.yml not provided", env)
		log.Panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)

	return conf
}
