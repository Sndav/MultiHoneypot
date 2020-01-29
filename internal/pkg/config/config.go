package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	ini *ini.File
}

func NewConfig(configFile string) *Config{
	conf := &Config{}
	c, err := ini.Load(configFile)
	if err != nil{
		panic(err)
	}
	conf.ini = c
	return conf
}

func (conf *Config) Get(node string, key string) string {
	val := conf.ini.Section(node).Key(key).String()
	return val
}

func (conf *Config) GetInt(node string, key string) int {
	val, _ := conf.ini.Section(node).Key(key).Int()
	return val
}