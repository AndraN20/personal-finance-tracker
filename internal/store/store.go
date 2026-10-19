// Package store keeps all of the service's state in memory.
package store

import "sync"

// Store is safe for concurrent use.
type Store struct {
	mu sync.Mutex
	// Transactions, rules and budgets will live here.
}

func New() *Store {
	return &Store{}
}

// Reset discards everything the store holds.
func (s *Store) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
}
