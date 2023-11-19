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
}

func (s *searcher) Search(_ context.Context, query string) ([]string, error) {
	normalizedQuery := strings.Fields(strings.ToLower(query))
	var results []string
	for emoji, labels := range emojis {
		contains := true
		for _, query := range normalizedQuery {
			if !slices.Contains(labels, query) {
				contains = false
				break
			}
		}
		if contains {
			results = append(results, emoji)
		}
	}
	sort.Strings(results)
	return results, nil
}
