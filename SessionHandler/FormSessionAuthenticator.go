package SessionHandler

import "net/http"

// A FormSessionAuthenticator ties together all of the SessionHand
type FormSessionAuthenticator struct {
	// A session store to put validated sessions into
	Storage SessionStore
	// A TokenExtractor to use to identify the session
	// being validated to the SessionStore
	Tokener TokenExtractor

	// A callback function to use to compare the username
	// and password submitted and return true if it's valid,
	// or false otherwise. Users of a FormSessionAuthenticator
	// are expected to implement this and pass it in the
	// constructor. See See examples/SessionValidatorExample
	// for an example
	ComparisonCallback func(username, password string) bool
}

// Implementation of SessionAuthenticator interface.
// Extracts the username and password form elements, and
// delegates to the ComparisonCallback which the application
// provided, storing the session in the SessionStore if it's
// valid.
func (s FormSessionAuthenticator) Authenticate(r *http.Request) bool {
	if s.ComparisonCallback(r.FormValue("username"), r.FormValue("password")) {
		token := s.Tokener.GetToken(r)
		s.Storage.AddSession(token, Session{})
		return true
	}

	return false
}
