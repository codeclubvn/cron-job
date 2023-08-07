package conf

import (
	"cron-job/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	ConfigFile string `env:"CONFIG_FILE" envDefault:"config.yml"`
}

func NewConfig() *Config {
	return &Config{}
}

// LoadConfigJobs loads jobs from yaml file
func (c *Config) LoadConfigJobs() ([]*model.JobConfig, error) {
	raw, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("failed read config file: %v", err)
	}

	var jobs []*model.JobConfig
	err = yaml.Unmarshal(raw, &jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config data: %v", err)
	}

	return jobs, nil
}
