package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/apex/gateway/v2"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"

	"github.com/matsilva/goinstall/resolver"
	"github.com/matsilva/goinstall/server"
	gstorage "github.com/matsilva/goinstall/storage"
)

func main() {

	// context
	ctx := context.Background()

	// github client
	gh := os.Getenv("GITHUB_TOKEN")

	// storage client
	gs, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("error creating storage client: %s\n", err)
		os.Exit(1)
	}

	// server
	s := &server.Server{
		Resolver: &resolver.GitHub{
			Client: github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: gh},
			))),
		},
		Storage: gstorage.NewGoogle(&gstorage.GoogleOptions{
			Client: gs,
			Bucket: "goinstall",
		}),
	}

	// serve
	addr := ":" + os.Getenv("PORT")
	if os.Getenv("UP_STAGE") != "" {
		fmt.Printf("listening on lambda\n")
		gateway.ListenAndServe(addr, s)
		return
	}

	fmt.Printf("listening on %s\n", addr)
	if err := http.ListenAndServe(addr, s); err != nil {
		fmt.Printf("error serving: %s\n", err)
		os.Exit(1)
	}
}
