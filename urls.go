package main

import "gopkg.in/simversity/gottp.v1"

func init() {
	gottp.NewUrl("shorten", "/shorten/?$", new(ShortenerHandler))
	gottp.NewUrl("redirect", "/\\w{6,10}/?$", new(RedirectHandler))
}
