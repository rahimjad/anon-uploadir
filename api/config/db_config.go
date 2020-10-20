package config

import "fmt"

// DbConfig is a struct used to capture database configuration values from the yml file provided
type DbConfig struct {
	HOST     string `yaml:"host"`
	PORT     int64  `yaml:"port"`
	USER     string `yaml:"user"`
	PASSWORD string `yaml:"password"`
	DBNAME   string `yaml:"dbname"`
	SSLMODE  string `yaml:"sslmode"`
}

// BuildConnectionString generates the connection string from the DbConfig struct needed to connect to the DB
func (conf DbConfig) BuildConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.HOST,
		conf.PORT,
		conf.USER,
		conf.PASSWORD,
		conf.DBNAME,
		conf.SSLMODE)
}
