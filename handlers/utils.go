package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Simversity/shortener/db"
	"gopkg.in/simversity/gottp.v1"
)

var mutex sync.Mutex

type Redirect struct {
	gottp.BaseHandler
}

func (self *Redirect) Get(req *gottp.Request) {
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

	actual_entity := db.UrlModel{}

	err := db.GetOne("short_url", shortkey, &actual_entity)
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

func CreateLink(url_object *db.UrlModel) {
	err := db.GetOne("url", url_object.Url, url_object)
	if err == nil {
		return
	}

	if url_object.ShortUrl == "" {
		mutex.Lock()
		defer mutex.Unlock()

		short_url := createShortString()

		for true {
			if db.Count("short_url", short_url) == 0 {
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

type Shortener struct {
	gottp.BaseHandler
}

func (self *Shortener) Post(req *gottp.Request) {
	arg_url := (*req.GetArguments())["url"]

	if arg_url == nil {
		e := gottp.HttpError{412, "URL Not found"}
		req.Raise(e)
		return
	}

	a_url, _ := arg_url.(string)
	if a_url == "" {
		e := gottp.HttpError{412, "URL Not found"}
		req.Raise(e)
		return
	}

	parsed_url, parse_err := url.Parse(a_url)
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
		a_url = "http://" + a_url
	}

	url_object := db.UrlModel{Url: a_url}
	CreateLink(&url_object)

	var short_host string
	short_host_arg := (*req.GetArguments())["shortener_host"]
	if short_host_arg != nil {
		short_host, _ = short_host_arg.(string)
	}

	if short_host == "" {
		short_host = "http://" + req.Request.Host
	}

	url_object.ShortUrl = short_host + "/" + url_object.ShortUrl
	req.Write(url_object)
	return
}
