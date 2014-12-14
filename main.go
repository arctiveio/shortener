// +build !appengine

package main

import (
	"github.com/Simversity/gottp"
	"github.com/Simversity/shortener/db"
)

func sysInit() {
	<-(gottp.SysInitChan)
	db.InitDB(Settings.Shortener.StoragePath)
}

func init() {
	gottp.BindHandlers(Urls)
}

func main() {
	go sysInit()
	gottp.MakeServer(&Settings)
}
