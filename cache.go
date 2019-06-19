package eurex

import (
	cache "github.com/patrickmn/go-cache"
)

// CacheService handles in-memory caching our rates
type CacheService struct {
	client *Client
	store  *cache.Cache
}

// NewCacheService creates a new handler for this service
func NewCacheService(client *Client, store *cache.Cache) *CacheService {
	return &CacheService{
		client,
		store,
	}
}

// Get will return our in-memory stored currency/rates per day
func (s *CacheService) Get(time string) (*map[string]float64, bool) {
	if x, found := s.store.Get(time); found {
		var result = x.(map[string]float64)
		return &result, found
	}
	return nil, false
}

// Store will store our currency/rates in-memory
func (s *CacheService) Store(response *ExchangeRate) {
	s.store.Set(
		response.Date,
		response.Rates,
		cache.DefaultExpiration,
	)
}

// IsExpired checks whether or not currency/rates stored is expired
func (s *CacheService) IsExpired(base string) bool {
	if _, found := s.store.Get(base); found {
		return false
	}
	return true
}

// Expire will expire the cache for a given date
func (s *CacheService) Expire(time string) {
	s.store.Delete(time)
}
