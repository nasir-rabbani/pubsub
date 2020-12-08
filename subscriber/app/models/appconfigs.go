package models

var (

	// AppConfig - to store application specific configurations
	AppConfig appConfig

	// ConfigFilePath - Path to application related configs
	ConfigFilePath = "./configs/appconfig.yaml"

	// DbConfigPath - Path to database related configs
	DbConfigPath = "./configs/dbconfig.yaml"
)

// appConfig - appConfig
type appConfig struct {
	Port       int        `yaml:"port"`
	AmqpConfig amqpConfig `yaml:"amqpConfig"`
}

type amqpConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
