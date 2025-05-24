package main

import (
	"context"
	"fmt"
	"os"
)

var (
	defaultGPT *catGPT
)

func main() {
	fmt.Println("hi!")
	listenPublic := ":8080"
	if lp := os.Getenv("CATGPT_LISTEN_PUBLIC"); lp != "" {
		listenPublic = lp
	}

	// This listener should not be exposed to the internet.
	listenPrivate := ":9090"
	if lp := os.Getenv("CATGPT_LISTEN_PRIVATE"); lp != "" {
		listenPrivate = lp
	}

	defaultGPT = &catGPT{}
	serve(context.Background(), listenPublic, listenPrivate)
}
