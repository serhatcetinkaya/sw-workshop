package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

// Cache caches query results.
type Cache interface {
	// Get returns the cached emojis produced by the provided query. On cache
	// miss, Get returns nil, nil.
	Get(context.Context, string) ([]string, error)

	// Put stores a query and its corresponding emojis in the cache.
	Put(context.Context, string, []string) error
}

type cacher struct {
	weaver.Implements[Cache]
	fakeCache map[string][]string
}

func (c *cacher) Init(context.Context) error {
	c.fakeCache = map[string][]string{}
	return nil
}

func (c *cacher) Get(ctx context.Context, query string) ([]string, error) {
	return c.fakeCache[query], nil
}

func (c *cacher) Put(ctx context.Context, query string, result []string) error {
	c.fakeCache[query] = result
	fmt.Println(c.fakeCache)
	return nil
}
