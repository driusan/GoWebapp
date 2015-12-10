package SessionHandler

import "net/http"

// A session of a client to the web application.
// This is expected to eventually support a serializable
// key/value map
type Session struct {
}

// A Token represents a way to map a request to a session.
// It can be extracted from an http.Request object by anything
// that implements the TokenExtractor interface
type Token string

// An interface to something that stores sessions.
// MemorySessionHandler is the only currently implemented
// SessionStore, but more persistent stores (ie. a
// DatabaseSessionStore, or a FilesystemSessionStore)
// should support this interface.
type SessionStore interface {
	GetSession(Token) *Session
	AddSession(Token, Session) error
}

// A SessionAuthenticator takes an http.Request, and
// validates whether or not it is authenticated, and
// should store the Token of an authenticated token
// into a SessionStore.
//
// See FormSessionAuthenticator for an example which
// validates Sessions based on the HTML form submission
//
// In the future, new SessionAuthenticators such as a
// BasicAuth authenticator or a Authorization: Bearer
// (Oauth, JWT) token authenticator, or a cookie
// authenticator should also be implemented.
type SessionAuthenticator interface {
	Authenticate(r *http.Request) bool
}

// A TokenExtractor is something that takes a HTTP
// request object, and extracts a token which can be
// used by a SessionStore to identify the session
// making the request.
//
// See the SessionValidatorExample application for
// an example implementation.
//
// In the future, new TokenExtractors such as a
// an Authorization: Bearer token extractor,
// or a cookie token extractor should also be supported.
type TokenExtractor interface {
	GetToken(r *http.Request) Token
}
