// +build !appengine

package main

import (
	"gopkg.in/simversity/gottp.v1"
	"gopkg.in/simversity/shortener.v1/db"
)

func sysInit() {
	<-(gottp.SysInitChan)
	db.InitDB(Settings.Shortener.StoragePath)
}

func main() {
	go sysInit()
	gottp.MakeServer(&Settings)
}
