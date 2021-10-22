package cache

import (
	gocache "github.com/patrickmn/go-cache"
)

type Settings struct {
	ExpiryInSeconds  int `mapstructure:"expiry_in_seconds"`
	CleanupInSeconds int `mapstructure:"cleanup_in_seconds"`
}
type CacheClient interface {
}

type cacheClient struct {
	gocache gocache.Cache
}

func NewCacheClient() {
	c := gocache.New(gocache.DefaultExpiration)
}
