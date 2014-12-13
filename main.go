// +build !appengine

package main

import (
	"shortener/db"

	"github.com/Simversity/gottp"
)

func sysInit() {
	<-(gottp.SysInitChan)
	db.InitDB(Settings.Shortener.StoragePath)
}

func main() {
	go sysInit()
	gottp.MakeServer(&Settings)
}
