package syncache

import (
	"sync"
	"time"
)

type entry struct {
	value interface{}
	ttl   time.Time
}

func NewWithExpireFunc(isExpired func(t time.Time) bool) *Storage {
	return &Storage{
		table:     make(map[string]entry),
		isExpired: isExpired,
	}
}

func NewWithDefaultExpireFunc() *Storage {
	return NewWithExpireFunc(func(t time.Time) bool {
		return time.Since(t) > 0
	})
}

type Storage struct {
	table     map[string]entry
	mut       sync.RWMutex
	isExpired func(time.Time) bool
}

func (s *Storage) get(key string) (interface{}, bool) {
	item, ok := s.table[key]
	if !ok || s.isExpired(item.ttl) {
		return nil, false
	}
	return item.value, true
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	return s.get(key)
}

func (s *Storage) Set(key string, value interface{}, ttl time.Time, respect bool) {
	if value == nil || s.isExpired(ttl) {
		return
	}

	s.mut.Lock()
	if respect {
		_, respect = s.get(key)
	}
	if !respect {
		s.table[key] = entry{
			value: value,
			ttl:   ttl,
		}
	}
	s.mut.Unlock()
}

func (s *Storage) Delete(key string) {
	s.mut.Lock()
	delete(s.table, key)
	s.mut.Unlock()
}

func (s *Storage) DeleteAll() {
	s.mut.Lock()
	s.table = make(map[string]entry)
	s.mut.Unlock()
}

func (s *Storage) DeleteExpired() {
	s.mut.Lock()
	for k, v := range s.table {
		if s.isExpired(v.ttl) {
			delete(s.table, k)
		}
	}
	s.mut.Unlock()
}

func (s *Storage) Len() int {
	s.mut.RLock()
	defer s.mut.RUnlock()

	return len(s.table)
}
