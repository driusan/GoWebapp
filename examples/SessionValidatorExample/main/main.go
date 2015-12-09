package main

import (
	"URLHandler"
	"SessionHandler"
	"net/http"
)

var validator SessionHandler.SessionAuthenticator;
// Implement a new type which inherits the DefaultHandler
// behaviour
type SimpleGetPage struct {
	URLHandler.DefaultHandler
}

type SimpleTokenExtractor struct {}

func (s SimpleTokenExtractor) GetToken(r *http.Request) SessionHandler.Token {
	return SessionHandler.Token(r.FormValue("password"))
}
// Override Get to return some text for the root page, and
// a 404 error for any other request
func (r SimpleGetPage) Get(req *http.Request, params map[string]interface{}) (string, error) {
	if validator.Authenticate(req) == false {
		return "You are not allowed here",URLHandler.ForbiddenError{}
	}
	if req.URL.Path == "/" {
		return "I am a page", nil
	}
	return "Page not found", URLHandler.NotFoundError{}

}

func formComparisonCallback(username, password string) (bool) {
	if password == "hunter2" {
		return true
	}
	return false
}
func main() {
	handle := SimpleGetPage{}

	storage := SessionHandler.MemorySessionStorer{
		Token: SimpleTokenExtractor{},
	}

	validator = SessionHandler.FormSessionAuthenticator{storage,
		formComparisonCallback}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
