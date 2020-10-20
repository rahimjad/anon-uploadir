package config

// AwsConfig is a struct used to capture aws configuration values from the yml file provided
type AwsConfig struct {
	Region          string `yaml:"region"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	BucketName      string `yaml:"bucket_name"`
}
