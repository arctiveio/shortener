GOPATH := $(shell cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)

deps: src
	#Dependencies
	GOPATH=$(GOPATH) go get code.google.com/p/gcfg
	GOPATH=$(GOPATH) go get github.com/HouzuoGuo/tiedot
	GOPATH=$(GOPATH) go get github.com/Simversity/gottp

install: deps
	# binary
	GOPATH=$(GOPATH) go install -a shortener

build: deps
	GOPATH=$(GOPATH) go build shortener

build-dry: deps
	# binary
	GOPATH=$(GOPATH) go build -n shortener
