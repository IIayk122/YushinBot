package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Configs ...
var Configs *Config

//Config ...
type Config struct {
	Bot   *BotConfig   `yaml:"bot"`
	Mongo *MongoConfig `yaml:"mongo"`
}

//Load ...
func Load(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	Configs = &Config{}
	err = yaml.Unmarshal(f, &Configs)
	if err != nil {
		return err
	}
	return nil
}

//BotConfig ...
type BotConfig struct {
	Token         string `yaml:"token"`
	PollerTimeout uint64 `yaml:"poller-timeout"`
}

//MongoConfig ...
type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Auth     string `yaml:"auth_source"`
}
