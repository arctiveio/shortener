package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/simversity/gottp.v2"
	"gopkg.in/simversity/gottp.v2/utils"
)

var mutex sync.Mutex

func ConcatenateErrors(errs *[]error) string {
	var errString string
	for i := 0; i < len(*errs); i++ {
		errString += (*errs)[i].Error()
		if (len(*errs) - i) > 1 {
			errString += ","
		}
	}
	return errString
}

type RedirectHandler struct {
	gottp.BaseHandler
}

func (self *RedirectHandler) Get(req *gottp.Request) {
	shortkey := req.Request.URL.String()

	if strings.HasPrefix(shortkey, "/") {
		shortkey = strings.TrimLeft(shortkey, "/")
	}

	if strings.HasPrefix(shortkey, "redirect") {
		shortkey = strings.TrimLeft(shortkey, "redirect")
	}

	if strings.HasPrefix(shortkey, "/") {
		shortkey = strings.TrimLeft(shortkey, "/")
	}

	actual_entity := UrlModel{}

	err := GetOne("short_url", shortkey, &actual_entity)
	if err != nil {
		e := gottp.HttpError{404, "/" + shortkey + " not Found"}
		req.Raise(e)
		return
	}

	parsed_url, parse_err := url.Parse(actual_entity.Url)

	if parse_err != nil {
		panic(parse_err)
	}

	q := parsed_url.Query()
	q.Set("sh", "1")
	parsed_url.RawQuery = q.Encode()
	http.Redirect(req.Writer, req.Request, parsed_url.String(), http.StatusMovedPermanently)
	return
}

func createShortString() string {
	t1 := int64(time.Now().Unix())
	return strconv.FormatInt(t1, 36)
}

func CreateLink(url_object *UrlModel) {
	err := GetOne("url", url_object.Url, url_object)
	if err == nil {
		return
	}

	if url_object.ShortUrl == "" {
		mutex.Lock()
		defer mutex.Unlock()

		short_url := createShortString()

		for true {
			if Count("short_url", short_url) == 0 {
				break
			} else {
				short_url = createShortString()
			}
		}

		url_object.ShortUrl = short_url
	}

	if err := url_object.Insert(); err != nil {
		panic(err)
	}
}

type ShortenerHandler struct {
	gottp.BaseHandler
}

func (self *ShortenerHandler) Post(req *gottp.Request) {
	shortenRequest := new(UrlModel)
	req.ConvertArguments(shortenRequest)

	errors := utils.Validate(shortenRequest)
	if len(*errors) > 0 {
		e := gottp.HttpError{http.StatusNotImplemented, ConcatenateErrors(errors)}
		req.Raise(e)
		return
	}

	parsed_url, parse_err := url.Parse(shortenRequest.Url)
	if parse_err != nil {
		e := gottp.HttpError{412, parse_err.Error()}
		req.Raise(e)
		return
	}

	if len(parsed_url.Host) < 1 {
		e := gottp.HttpError{412, "Cannot shorten Relative URLs. Invalid Domain: " + parsed_url.Host}
		req.Raise(e)
		return
	}

	if !strings.HasPrefix(parsed_url.Scheme, "http") {
		shortenRequest.Url = "http://" + shortenRequest.Url
	}

	CreateLink(shortenRequest)

	var short_host string
	short_host_arg := (*req.GetArguments())["shortener_host"]
	if short_host_arg != nil {
		short_host, _ = short_host_arg.(string)
	}

	if short_host == "" {
		short_host = "http://" + req.Request.Host
	}

	shortenRequest.ShortUrl = short_host + "/" + shortenRequest.ShortUrl
	req.Write(shortenRequest)
	return
}
