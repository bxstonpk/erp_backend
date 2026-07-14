package usecase

import (
	"sync"
	"time"
)

// revokedStore tracks refresh-token jtis revoked via logout or rotation.
// In-memory only: there is no session/token table yet, so revocation does
// not survive a restart or span multiple instances. Access tokens are not
// tracked here — they stay short-lived and expire naturally.
type revokedStore struct {
	mu    sync.Mutex
	items map[string]time.Time
}

func newRevokedStore() *revokedStore {
	return &revokedStore{items: make(map[string]time.Time)}
}

func (s *revokedStore) revoke(jti string, expiresAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cleanupLocked()
	s.items[jti] = expiresAt
}

func (s *revokedStore) isRevoked(jti string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.items[jti]
	return ok
}

func (s *revokedStore) cleanupLocked() {
	now := time.Now()
	for jti, expiresAt := range s.items {
		if now.After(expiresAt) {
			delete(s.items, jti)
		}
	}
}
