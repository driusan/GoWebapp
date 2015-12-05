package URLHandler

import (
	"fmt"
	"net/http"
)

// A map of extra things to pass to every request handler call
var extras map[string]interface{}

// URLHandler is an interface to describe a request to a URL
//
// After being registered to handle a URL with a RegisterHandler
// call, the URLHandler will handle any requests to that URL by
// delegating to the method for the appropriate HTTP Method being
// called.

// All methods receive the http.Request object, and a map of extra
// parameters that have been registered with RegisterExtraParameter
type URLHandler interface {
	// Get will handle an HTTP GET request to this URL and return the
	// content that should be sent to the client
	Get(r *http.Request, params map[string]interface{}) (string, error)

	// Post will handle an HTTP POST request to this URL.
	// Post returns 2 strings: the content to return, an a redirectURL
	// If the redirectURL is not the empty string, the registered
	// URLandler will automatically respond with a 302 return code
	// instead of a 200 return code, and set an appropriate Location:
	// response header
	Post(r *http.Request, params map[string]interface{}) (content, redirectURL string, err error)

	// Put will handle an HTTP PUT request to this URL and return the
	// content that should be sent to the client
	Put(r *http.Request, params map[string]interface{}) (string, error)

	// Delete will handle an HTTP PUT request to this URL and return the
	// content that should be sent to the client
	Delete(r *http.Request, params map[string]interface{}) (string, error)
}

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

// RegisterHandler takes a URLHandler and a url string and registers
// that URLHandler to handle that URL. It automatically registers an
// http.HandleFunc which delegates to the appropriate URLHandler method
func RegisterHandler(h URLHandler, url string) {
	var handler = func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			response, err := h.Get(r, extras)
			if err != nil {
				handled := handleClientError(w, response, err)
				if handled {
					return
				}
				panic("I got an error I didn't understand")

			}
			fmt.Fprintf(w, response)
		}
		if r.Method == "POST" {
			response, redirectURL, err := h.Post(r, extras)
			if err != nil {
				handled := handleClientError(w, response, err)
				if handled {
					return
				}
				panic("I got an error I didn't understand")
			}
			if redirectURL != "" {
				w.Header().Add("Location", redirectURL)
				w.WriteHeader(303)
			}
			fmt.Fprintf(w, response)
		}

	}
	http.HandleFunc(url, handler)
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