package main

import (
	"shortener/handlers"

	"github.com/Simversity/gottp"
)

var Urls = []*gottp.Url{
	gottp.NewUrl("shorten", "/shorten/?$", handlers.Shortener),
	gottp.NewUrl("redirect", "/\\w{6,10}/?$", handlers.Redirect),
}
