// +build !appengine

package main

import (
	"shortener/conf"
	"shortener/db"

	"github.com/Simversity/gottp"
)

func sysInit() {
	<-(gottp.SysInitChan)
	db.InitDB(conf.Settings.Shortener.StoragePath)
}

func main() {
	go sysInit()
	gottp.MakeServer(&conf.Settings)
}
