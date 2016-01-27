package URLHandler

import (
	"net/http"
	"net/url"
	"time"
)

var UnknownMTime time.Time = time.Time{}

// DefaultHandler is an simple implementation of the URLHandler
// interface that you can compose into your class if you only
// want to implement some methods.
// The DefaultHandler will respond with a 405 Method Not Allowed
// response to every request.
type DefaultHandler struct{}

func (url DefaultHandler) Get(r *http.Request, params map[string]interface{}) (string, error) {
	return "Method not implemented", InvalidMethodError{}
}
func (url DefaultHandler) Post(r *http.Request, params map[string]interface{}) (string, string, error) {
	return "Method not implemented", "", InvalidMethodError{}
}
func (url DefaultHandler) Put(r *http.Request, params map[string]interface{}) (string, error) {
	return "Method not implemented", InvalidMethodError{}
}
func (url DefaultHandler) Delete(r *http.Request, params map[string]interface{}) (string, error) {
	return "Method not implemented", InvalidMethodError{}
}

func (url DefaultHandler) ETag(r *url.URL, params map[string]interface{}) (ETag, error) {
	return "", nil
}

func (url DefaultHandler) LastModified(r *url.URL, params map[string]interface{}) time.Time {
	return UnknownMTime
}
