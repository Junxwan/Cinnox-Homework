package cmd

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
	"time"
)

var Conf *Config

type Config struct {
	Http      Http
	Line      Line
	Databases Databases
}

type Http struct {
	Addr  string
	Debug bool
}

type Line struct {
	Secret string
	Token  string
}

type Databases struct {
	Connections string
	Host        string
	Port        int
	Databases   string
	Username    string
	Password    string
	Timeout     time.Duration
}

func NewConfig(path string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CINNOX-HOMEWORK")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")

	if path == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		path = viper.ConfigFileUsed()
	}

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	if err = viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return nil, err
	}

	Conf = new(Config)
	return Conf, viper.Unmarshal(Conf)
}
