package main

import (
	"github.com/Simversity/gottp"
	"github.com/Simversity/shortener/handlers"
)

func init() {
	gottp.NewUrl("shorten", "/shorten/?$", new(handlers.Shortener))
	gottp.NewUrl("redirect", "/\\w{6,10}/?$", new(handlers.Redirect))
}
