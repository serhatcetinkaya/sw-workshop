package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
	searcher weaver.Ref[Searcher]
	listener weaver.Listener `weaver:"emojis"`
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, a *app) error {
	fmt.Printf("the app is listening: %s", a.listener)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		emojis, err := a.searcher.Get().Search(ctx, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(emojis)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(bytes))
	})
	return http.Serve(a.listener, nil)
}
