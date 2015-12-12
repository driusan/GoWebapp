package URLHandler

import (
	"fmt"
	"net/http"
	"net/url"
)

type ETag string

// A map of extra things to pass to every request handler call
var extras map[string]interface{}

// URLHandler is an interface to describe a request to a URL
//
// After being registered to handle a URL with a RegisterHandler
// call, the URLHandler will handle any requests to that URL by
// delegating to the method for the appropriate HTTP Method being
// called.
//
// All methods receive the http.Request object, and a map of extra
// parameters that have been registered with RegisterExtraParameter
type URLHandler interface {
	// Get will handle an HTTP GET request to this URL and return the
	// content that should be sent to the client
	Get(r *http.Request, params map[string]interface{}) (string, error)

	// Post will handle an HTTP POST request to this URL.
	// Post returns 2 strings: the content to return, an a redirectURL
	// If the redirectURL is not the empty string, the registered
	// URLandler will automatically respond with a 303 return code
	// instead of a 200 return code, and set an appropriate Location:
	// response header
	Post(r *http.Request, params map[string]interface{}) (content, redirectURL string, err error)

	// Put will handle an HTTP PUT request to this URL and return the
	// content that should be sent to the client
	Put(r *http.Request, params map[string]interface{}) (string, error)

	// Delete will handle an HTTP PUT request to this URL and return the
	// content that should be sent to the client
	Delete(r *http.Request, params map[string]interface{}) (string, error)

	// Calculate an ETag to represent the resource being served by
	// this handler, so that a registered handler can return a 304
	// code if the resource hasn't changed.
	ETag(*url.URL, map[string]interface{}) ETag
}

// handleClientError takes an error from a URLHandler and returns
// an appropriate response if it knows how. Returns true if it's been
// handled, false otherwise
func handleClientError(w http.ResponseWriter, response string, err error) bool {
	switch err.(type) {
	case ForbiddenError:
		w.WriteHeader(403)
		fmt.Fprintf(w, response)
		return true
	case NotFoundError:
		w.WriteHeader(404)
		fmt.Fprintf(w, response)
		return true
	case InvalidMethodError:
		w.WriteHeader(405)
		fmt.Fprintf(w, response)
		return true
	}
	return false
}

// handleError if a helper function to handle errors from URLHandlers.
// Mostly, it calls handleClientError and then panics if it didn't get
// handled.
func handleError(w http.ResponseWriter, response string, err error) {
	handled := handleClientError(w, response, err)
	if handled {
		return
	}
	panic("Something happened")
}

// RegisterHandler takes a URLHandler and a url string and registers
// that URLHandler to handle that URL. It automatically registers an
// http.HandleFunc which delegates to the appropriate URLHandler method
func RegisterHandler(h URLHandler, url string) {
	var handler = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unknown server error")
			}
		}()

		switch r.Method {
		case "GET":
			if etag := h.ETag(r.URL, extras); etag != "" {
				w.Header().Add("ETag", string(etag))
				if string(etag) == r.Header.Get("If-None-Match") {
					w.WriteHeader(304)
					return
				}
			}
			response, err := h.Get(r, extras)
			if err != nil {
				handleError(w, response, err)
				return
			}
			fmt.Fprintf(w, response)
		case "POST":
			response, redirectURL, err := h.Post(r, extras)
			if err != nil {
				handleError(w, response, err)
				return
			}
			if redirectURL != "" {
				w.Header().Add("Location", redirectURL)
				w.WriteHeader(303)
			}
			fmt.Fprintf(w, response)
		default:
			w.WriteHeader(501)

		}
	}
	http.HandleFunc(url, handler)
}

// RegisterStaticHandler registers directory to be served by the web
// server on the filesystem without going through the handler function.
func RegisterStaticHandler(url, directoryRoot string) {
	http.Handle(url, http.StripPrefix(url, http.FileServer(http.Dir(directoryRoot))))
}

// RegisterExtraParameter allows you to add arbitrary data to get
// passed to the params parameter of URLHandler handler functions which
// you can retrieve from params[key] (and will need to manually cast to
// the appropriate type.
//
// This is useful for passing, for instance, a pointer to an sql.DB,
// or any configuration data you want to use throughout your web app
func RegisterExtraParameter(key string, obj interface{}) {
	if extras == nil {
		extras = make(map[string]interface{})
	}
	extras[key] = obj
}
