package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Config Conf

type Environment struct {
	Development Conf `yaml:"development"`
	Production Conf `yaml:"production"`
}

type Conf struct {
	Database Database `yaml:"db"`
}

type Database struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
}

func SetEnvironment(env string) {

	buf, err := ioutil.ReadFile("./config/environment.yml")
	if err != nil {
		panic(err)
	}

	var environment Environment

	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		panic(err)
	}

	if env == "development" {
		Config = environment.Development
	} else {
		Config = environment.Production
	}
}
