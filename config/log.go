package config

type log struct {
	Name  string   `yaml:"name"`
	Path  string   `yaml:"path"`
	Split logSplit `yaml:"split"`
}

type logSplit struct {
	MaxSize    int  `yaml:"maxSize"`
	MaxAge     int  `yaml:"maxAge"`
	MaxBackups int  `yaml:"maxBackups"`
	Compress   bool `yaml:"compress"`
}
