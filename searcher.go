package main

import (
	"context"
	"slices"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver"
)

// Searcher component.
type Searcher interface {
	Search(ctx context.Context, query string) ([]string, error)
}

// implementation of the searcher component
type searcher struct {
	weaver.Implements[Searcher]
	cache weaver.Ref[Cache]
}

func (s *searcher) Search(ctx context.Context, query string) ([]string, error) {
	res, _ := s.cache.Get().Get(ctx, query)
	if res != nil {
		return res, nil
	}

	normalizedQuery := strings.Fields(strings.ToLower(query))
	results := []string{}
	for emoji, labels := range emojis {
		contains := true
		for _, q := range normalizedQuery {
			if !slices.Contains(labels, q) {
				contains = false
				break
			}
		}
		if contains {
			results = append(results, emoji)
		}
	}
	sort.Strings(results)
	_ = s.cache.Get().Put(ctx, query, results)
	return results, nil
}
