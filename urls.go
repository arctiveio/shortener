package main

import (
	"gopkg.in/simversity/gottp.v1"
	"gopkg.in/simversity/shortener.v1/handlers"
)

func init() {
	gottp.NewUrl("shorten", "/shorten/?$", new(handlers.Shortener))
	gottp.NewUrl("redirect", "/\\w{6,10}/?$", new(handlers.Redirect))
}
