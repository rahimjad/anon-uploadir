package config

// AwsConfig is a struct used to capture aws configuration values from the yml file provided
type RouterConfig struct {
	Port int `yaml:"port"`
}
