package main

import conf "github.com/Simversity/gottp/conf"

const baseConfig = `;Sample Configuration File
[gottp]
listen="/tmp/shortener.sock"
EmailHost=""
EmailPort=""
EmailPassword=""
EmailUsername=""
EmailSender="Shortener"
EmailFrom=""
ErrorTo=""
EmailDummy=true

[shortener]
StoragePath="/tmp/shortener" #Override this config file`

type config struct {
	Gottp     conf.GottpSettings
	Shortener struct {
		StoragePath string
	}
}

func (self *config) MakeConfig(configPath string) {
	conf.ReadConfig(baseConfig, self)
	if configPath != "" {
		conf.MakeConfig(configPath, self)
	}
}

func (self *config) GetGottpConfig() *conf.GottpSettings {
	return &self.Gottp
}

var Settings config
