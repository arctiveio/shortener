package main

import (
	"github.com/Simversity/gottp"
	"github.com/Simversity/shortener/handlers"
)

var Urls = []*gottp.Url{
	gottp.NewUrl("shorten", "/shorten/?$", handlers.Shortener),
	gottp.NewUrl("redirect", "/\\w{6,10}/?$", handlers.Redirect),
}
