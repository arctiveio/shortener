package main

import conf "gopkg.in/simversity/gottp.v2/conf"

const baseConfig = `;Sample Configuration File
[gottp]
listen="0.0.0.0:9010"
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
