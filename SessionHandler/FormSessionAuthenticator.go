package SessionHandler

import "net/http"

type FormSessionAuthenticator struct {
	Storage            SessionStorer
	ComparisonCallback func(username, password string) bool
}

func (s FormSessionAuthenticator) Authenticate(r *http.Request) bool {
	if s.ComparisonCallback(r.FormValue("username"), r.FormValue("password")) {
		s.Storage.AddSession(Session{r})
		return true
	}

	return false
}
