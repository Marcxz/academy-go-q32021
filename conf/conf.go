package conf

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server            string `yaml:"server"`
	Base_path         string `yaml:"base_path"`
	Filename          string `yaml:"filename"`
	Api_url           string `yaml:"api_url"`
	Postgres_host     string `yaml:"postgres_host"`
	Postgres_port     string `yaml:"postgres_post"`
	Postgres_db       string `yaml:"postgres_db"`
	Postgres_user     string `yaml:"postgres_user"`
	Postgres_password string `yaml:"postgres_password"`
}

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

func ParseFlags() (string, error) {
	var cp string

	flag.StringVar(&cp, "conf", "conf/conf.yml", "path config file")

	flag.Parse()

	if err := ValidateConfigPath(cp); err != nil {
		return "", err
	}

	return cp, nil
}
