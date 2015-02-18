// +build !appengine

package main

import "gopkg.in/simversity/gottp.v2"

func sysInit() {
	<-(gottp.SysInitChan)
	InitDB(Settings.Shortener.StoragePath)
}

func main() {
	go sysInit()
	gottp.MakeServer(&Settings)
}
