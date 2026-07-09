package store

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sync"
)

const maxCodeGenAttempts = 10
const slugLength = 7

type Store struct {
	codesMap   map[string]string
	reverseMap map[string]string
	mu         sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		codesMap:   make(map[string]string),
		reverseMap: make(map[string]string),
	}
}

func (s *Store) Get(slug string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.codesMap[slug]
	return value, ok
}

func (s *Store) Set(url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if value, ok := s.reverseMap[url]; ok {
		return value, nil
	}

	for i := 0; i < maxCodeGenAttempts; i++ {
		slug, err := generateSlug()
		if err != nil {
			return "", err
		}

		if _, ok := s.codesMap[slug]; !ok {
			s.codesMap[slug] = url
			s.reverseMap[url] = slug
			return slug, nil
		}
	}
	return "", errors.New("failed to generate unique code after max attempts")
}

func generateSlug() (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slug := make([]byte, slugLength)
	for i := range slug {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		slug[i] = charset[num.Int64()]
	}
	return string(slug), nil
}
