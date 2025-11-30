package storage

import "sync"

// MemoryStore is our in-memory database
// data is a map of string that is short code : url
type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemoryStore initializes and return a new MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string), // creates an empty map
	}
}

// Save stores a new mapping shortCode : Url
func (m *MemoryStore) Save(code, url string) {
	m.mu.Lock()         // lock for writing to prevent concurrent operations while writing
	defer m.mu.Unlock() // Unlock after function complete normally or accidentally

	m.data[code] = url
}

// Get retrieves the URL from the shortcode in the map
// It returns (url, true) or ("", false)
func (m *MemoryStore) Get(code string) (string, bool) {
	m.mu.RLock()         // to prevent any writing option read lock is applied, it will allow reading
	defer m.mu.RUnlock() // Unlock after func ends

	url, ok := m.data[code] // ok is true if key exists
	return url, ok
}
