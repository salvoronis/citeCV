package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

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
	JwtSecret	string `yaml:"secret"`
	RedisAddr	string `yaml:"redisaddr"`
}

func init() {
	conf = NewConfig("./configs.yaml")
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

func GetRoot() string {
	return conf.Host+conf.Port
}

func GetDbConnStr() string {
	tmpl, err := template.New("DBConnect").Parse("postgres://{{.Dbuser}}:{{.Dbpassword}}@{{.Dbhost}}:{{.Dbport}}/{{.Dbname}}?sslmode=disable")
	if err != nil {
		log.Printf("Can't parse db connection string %v\n", err)
	}
	var tmp bytes.Buffer
	if err := tmpl.Execute(&tmp, conf); err != nil {
		log.Printf("Can't execute configt into template %v\n", err)
	}
	return tmp.String()
}

func GetSecret() string {
	return conf.JwtSecret
}

func GetRedisAddr() string {
	return conf.RedisAddr
}
