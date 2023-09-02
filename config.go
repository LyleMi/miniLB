package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Entry struct {
	ListenAddress string `yaml:"listen"`
	TargetService string `yaml:"service"`
}

type Config struct {
	Entries []Entry `yaml:"entries"`
}

func ReadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
