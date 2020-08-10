package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Sshrw  string `yaml:"sshrw"`
	Config struct {
		Tunnel       string   `yaml:"tunnel"`
		Password     string   `yaml:"password"`
		Destinations []string `yaml:"destinations"`
	} `yaml:"config"`
}

func GetConfig() Config {
	t := Config{}
	buffer, err := ioutil.ReadFile("ssh_rw.yaml")
	err = yaml.Unmarshal(buffer, &t)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return t;
}
