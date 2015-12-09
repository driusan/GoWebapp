package SessionHandler

import "net/http"

type MemorySessionStorer struct {
	Token    TokenExtractor
	sessions map[Token]*Session
}

func (s MemorySessionStorer) AddSession(ses Session) error {
	if s.sessions == nil {
		s.sessions = make(map[Token]*Session)
	}

	token := s.Token.GetToken(ses.request)
	s.sessions[token] = &ses
	return nil
}

func (s MemorySessionStorer) GetSession(r *http.Request) *Session {

	if s.sessions == nil {
		return nil
	}
	token := s.Token.GetToken(r)
	return s.sessions[token]
}
