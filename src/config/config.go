package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var (
	conf Config
)

type Config struct {
	Host		string `yaml:"host"`
	Port		string `yaml:"port"`
	Dbhost		string `yaml:"dbhost"`
	Dbport		string `yaml:"dbport"`
	Dbuser		string `yaml:"dbuser"`
	Dbpassword	string `yaml:"dbpassword"`
	Dbname		string `yaml:"dbname"`
	Email		string `yaml:"email"`
	Pass		string `yaml:"pass"`
	Smtp		string `yaml:"smtp"`
	Smtpport	string `yaml:"smtpport"`
}

func init() {
	conf = NewConfig("../configs.yaml")
}

func NewConfig(path string) (conf Config) {
	confFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Can not read config file %v\n", err)
	}
	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		log.Printf("Config unmarshal error %v\n", err)
	}
	return conf
}

func GetRoot() string{
	return conf.Host+conf.Port
}
