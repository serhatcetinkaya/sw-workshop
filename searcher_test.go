package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/go-cmp/cmp"
)

// func TestSearch(t *testing.T) {
// 	runner := weavertest.Local // A runner that runs components in a single process
// 	runner.Test(t, func(t *testing.T, searcher Searcher) {
// 		ctx := context.Background()
// 		got, err := searcher.Search(ctx, "pig")
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		want := []string{"🐖", "🐗", "🐷", "🐽"}
// 		diff := cmp.Diff(got, want)

// 		if diff != "" {
// 			t.Fatalf("got: %s, diff: %s", got, diff)
// 		}
// 	})
// }

// proposed solution from docs
func TestSearch(t *testing.T) {
	type test struct {
		query string
		want  []string
	}

	for _, test := range []test{
		{"pig", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"PiG", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"black cat", []string{"🐈\u200d⬛"}},
		{"foo bar baz", []string{}},
	} {
		for _, runner := range weavertest.AllRunners() {
			runner.Name = fmt.Sprintf("%s/%q", runner.Name, test.query)
			runner.Test(t, func(t *testing.T, searcher Searcher) {
				got, err := searcher.Search(context.Background(), test.query)
				if err != nil {
					t.Fatalf("Search: %v", err)
				}
				if diff := cmp.Diff(test.want, got); diff != "" {
					t.Fatalf("Search (-want,+got):\n%s", diff)
				}
			})
		}
	}
} 