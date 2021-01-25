package databaseutils

import (
	"log"

	radix "github.com/mediocregopher/radix/v3"
)

var (
	redisCli *radix.Pool
)

func init() {
	var err error
	redisCli, err = radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		log.Println(err)
	}
}

func SetVal(key string, value string) error {
	return redisCli.Do(radix.Cmd(nil, "SET", key, value))
}

func GetVal(key string) (interface{}, error) {
	var value interface{}
	err := redisCli.Do(radix.Cmd(&value, "GET", key))
	return value, err
}
