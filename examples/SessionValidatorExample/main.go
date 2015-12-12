package main

//An application which includes a simple user validation/authentication.
//If the password "hunter2" is in the query string (or form
//submission), it will pass. Otherwise, it will return a 403 error.

import (
	"github.com/driusan/GoWebapp/SessionHandler"
	"github.com/driusan/GoWebapp/URLHandler"
	"net/http"
)

// Store a global variable with the current validator so that
// it's accessible. This could also be passed via RegisterExtraParameter
// to the URLHandler if you're allergic to global variables, but
// this way the compiler can enforce type information.
var validator SessionHandler.SessionAuthenticator

// Define a callback which the FormSessionAuthenticator
// can use to determine if this is a valid session. It
// ignores the username and just checks if the password
// is "hunter2"
func formComparisonCallback(username, password string) bool {
	if password == "hunter2" {
		return true
	}
	return false
}

// Define a new type of TokenExtractor
type SimpleTokenExtractor struct{}

// Extract a token to identify this session. Since only the
// password is used to identify it in this simple, we pretend
// it's a sufficient identification token. In real life, you'd
// want something (much) more secure than this to map requests
// to identification tokens such as a secure cookie with a unique
// hash.
func (s SimpleTokenExtractor) GetToken(r *http.Request) SessionHandler.Token {
	return SessionHandler.Token(r.FormValue("password"))
}

// This should look familiar from SimpleGetApp
type SimpleGetPage struct {
	URLHandler.DefaultHandler
}

func (r SimpleGetPage) Get(req *http.Request, params map[string]interface{}) (string, error) {
	// This is the only difference in this handler from SimpleGetApp.
	// Use the validator to validate this request, and return a 403
	// if the password is wrong.
	if validator.Authenticate(req) == false {
		return "You are not allowed here", URLHandler.ForbiddenError{}
	}
	if req.URL.Path == "/" {
		return "I am a page", nil
	}
	return "Page not found", URLHandler.NotFoundError{}

}

func main() {

	// Create a SessionStore
	storage := SessionHandler.MemorySessionStore{}

	// Create a FormSessionAuthenticator which uses the
	// above SessionStore, the SimpleTokenExtractor
	// and the formComparisonCallback to defined above
	// to validate.
	validator = SessionHandler.FormSessionAuthenticator{
		storage,
		SimpleTokenExtractor{},
		formComparisonCallback}

	// And finally, create the URLHandler and Register it.
	handle := SimpleGetPage{}
	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
