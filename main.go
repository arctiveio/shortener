// +build !appengine

package main

import (
	"gopkg.in/simversity/gottp.v1"
	"github.com/Simversity/shortener/db"
)

func sysInit() {
	<-(gottp.SysInitChan)
	db.InitDB(Settings.Shortener.StoragePath)
}

func main() {
	go sysInit()
	gottp.MakeServer(&Settings)
}
