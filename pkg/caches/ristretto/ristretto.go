package ristretto

import "github.com/dgraph-io/ristretto"

func NewLocalcahe() *ristretto.Cache {
	config := &ristretto.Config{
		NumCounters: 10 * 100_000_000,
		MaxCost:     100_000,
		BufferItems: 64,
		Metrics:     false,
		Cost: func(value interface{}) int64 {
			if v, ok := value.(string); ok {
				return int64(len(v))
			}
			return 1000
		},
	}
	cache, err := ristretto.NewCache(config)
	if err != nil {
		panic(err)
	}

	return cache
}
