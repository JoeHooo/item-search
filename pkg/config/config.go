package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server  Server  `json:"server"`
	Elastic Elastic `json:"elastic"`
	Redis   Redis   `json:"redis"`
	MySQL   MySQL   `json:"mysql"`
}

type Elastic struct {
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
	Port     int32  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Index    string `json:"index"`
}

type Server struct {
	Name         string `json:"name"`
	ContextPath  string `json:"contextPath"`
	Port         int32  `json:"port"`
	RunMode      string `json:"runMode"`
	ReadTimeOut  int32  `json:"readTimeOut"`
	WriteTimeOut int32  `json:"writeTimeOut"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var Conf Config

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config error, message: %s", err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Read config unmarshal error, message: %s", err)
	}
}
