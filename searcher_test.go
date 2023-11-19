package main

import (
	"context"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	runner := weavertest.Local // A runner that runs components in a single process
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		got, err := searcher.Search(ctx, "pig")
		if err != nil {
			t.Fatal(err)
		}

		want := []string{"🐖", "🐗", "🐷", "🐽"}
		diff := cmp.Diff(got, want)

		if diff != "" {
			t.Fatalf("got: %s, diff: %s", got, diff)
		}
	})
}