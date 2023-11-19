package main

import (
	"context"
	"fmt"
	"log"

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
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, a *app) error {
	fmt.Println("Hello, World!")
	emojis, err := a.searcher.Get().Search(ctx, "pig")
	if err != nil {
		return err
	}
	fmt.Println(emojis)
	return nil
}
