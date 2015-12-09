package SessionHandler

import "net/http"

type Session struct {
	request *http.Request
}

type Token string

type SessionStorer interface {
	GetSession(r *http.Request) *Session
	AddSession(Session) error
}

type SessionAuthenticator interface {
	Authenticate(r *http.Request) bool
}

type TokenExtractor interface {
	GetToken(r *http.Request) Token
}
