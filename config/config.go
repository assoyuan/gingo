package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type app struct {
	Name    string `yaml:"name"`
	Addr    string `yaml:"addr"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
}

type mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type logSplit struct {
	MaxSize    int  `yaml:"maxSize"`
	MaxAge     int  `yaml:"maxAge"`
	MaxBackups int  `yaml:"maxBackups"`
	Compress   bool `yaml:"compress"`
}

type log struct {
	Name  string   `yaml:"name"`
	Path  string   `yaml:"path"`
	Split logSplit `yaml:"split"`
}

type allConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	Log   log   `yaml:"log"`
}

var Config allConfig

func InitConfig() {
	baseDir, _ := os.Getwd()
	configFile := baseDir + "/config.yaml"

	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s\n", err))
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("配置文件解析失败: %s\n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&Config); err != nil {
			panic(fmt.Errorf("配置重载失败: %s\n", err))
		}
	})
}
