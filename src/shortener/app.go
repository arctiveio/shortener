package main

import (
	"github.com/Simversity/gottp"
)

func init() {
	gottp.BindHandlers(Urls)
}
