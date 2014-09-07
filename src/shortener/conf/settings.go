package conf

import (
	"github.com/Simversity/gottp"
	utils "github.com/Simversity/gottp/utils"
)

type config struct {
	Gottp     gottp.SettingsMap
	Shortener struct {
		StoragePath string
	}
}

func (self *config) MakeConfig(configPath string) {
	utils.ReadConfig(baseConfig, self)
	if configPath != "" {
		utils.MakeConfig(configPath, self)
	}
	gottp.Settings = self.Gottp
}

var Settings config
