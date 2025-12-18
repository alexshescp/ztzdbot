// internal/bot/session/store.go
package session

import "sync"

type Store struct {
	mu       sync.RWMutex
	sessions map[int64]*Session
}

func NewStore() *Store {
	return &Store{
		sessions: make(map[int64]*Session),
	}
}

func (s *Store) Get(chatID int64) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sess, ok := s.sessions[chatID]; ok {
		return sess
	}

	sess := &Session{
		ChatID: chatID,
		Step:   StepIdle,
	}
	s.sessions[chatID] = sess
	return sess
}

func (s *Store) Delete(chatID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, chatID)
}
