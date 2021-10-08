package conf

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config - the config struct for global variables
type Config struct {
	Server           string `yaml:"server"`
	BasePath         string `yaml:"basePath"`
	Filename         string `yaml:"filename"`
	ApiUrl           string `yaml:"apiURL"`
	MapPath          string `yaml:"mapPath"`
	MapFile          string `yaml:"mapFile"`
	PostgresHost     string `yaml:"postgresHost"`
	PostgresPort     string `yaml:"postgresPort"`
	PostgresDb       string `yaml:"postgresDB"`
	PostgresUser     string `yaml:"postgresUser"`
	PostgresPassword string `yaml:"postgresPassword"`
}

// NewConfi - Constructor function to create a config struct with the constant variables
func NewConfig(p string) (*Config, error) {
	c := &Config{}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	d := yaml.NewDecoder(f)

	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}

// ValidateConfigPath - Validate if the config file path exists and if it is a file
func ValidateConfigPath(p string) error {
	s, err := os.Stat(p)

	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, should be a file", p)
	}

	return nil
}

// ParseFlags - func to reference to a config path and validate it
func ParseFlags() (string, error) {
	var cp string

	flag.StringVar(&cp, "conf", "conf/conf.yml", "path config file")

	flag.Parse()

	if err := ValidateConfigPath(cp); err != nil {
		return "", err
	}

	return cp, nil
}
