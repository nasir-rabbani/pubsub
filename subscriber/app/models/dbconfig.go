package models

// DbConfigs - to map all the connection configs from dbconfig.yaml
type DbConfigs struct {
	Hosts map[string]DBConfiguration `yaml:"hosts"`
}

// DBConfiguration - to map basic attributes of DB connection
type DBConfiguration struct {
	Server    string `yaml:"server"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
	IsDefault bool   `yaml:"isDefault"`
}
