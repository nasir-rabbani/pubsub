package confighelper

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Init - Reads the given config.yaml file and maps configs
func Init(configPath string, config interface{}) error {
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(config); err != nil {
		return err
	}
	return nil
}
