GOPATH := $(shell cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)

install: src
	#Dependencies
	GOPATH=$(GOPATH) go get code.google.com/p/gcfg
	GOPATH=$(GOPATH) go get github.com/HouzuoGuo/tiedot
	GOPATH=$(GOPATH) go get github.com/Simversity/gottp
	# binary
	GOPATH=$(GOPATH) go install -a shortener

build-dry: src
	# binary
	GOPATH=$(GOPATH) go build -n shortener
