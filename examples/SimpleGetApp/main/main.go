package main

import (
	"URLHandler"
	"net/http"
)

// Implement a new type which inherits the DefaultHandler
// behaviour
type SimpleGetPage struct {
	URLHandler.DefaultHandler
}

// Override Get to return some text for the root page, and
// a 404 error for any other request
func (r SimpleGetPage) Get(req *http.Request, params map[string]interface{}) (string, error) {
	if req.URL.Path == "/" {
		return "I am a page", nil
	}
	return "Page not found", URLHandler.NotFoundError{}

}
func main() {
	handle := SimpleGetPage{}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
