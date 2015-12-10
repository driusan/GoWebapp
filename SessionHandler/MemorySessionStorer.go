package SessionHandler

// A MemorySessionStore provides the simplest SessionStore
// possible. It stores active Tokens to Sessions in memory.
type MemorySessionStore struct {
	sessions map[Token]*Session
}

func (s MemorySessionStore) AddSession(t Token, ses Session) error {
	if s.sessions == nil {
		s.sessions = make(map[Token]*Session)
	}

	s.sessions[t] = &ses
	return nil
}

func (s MemorySessionStore) GetSession(t Token) *Session {

	if s.sessions == nil {
		return nil
	}
	return s.sessions[t]
}
