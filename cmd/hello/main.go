package main

import (
	"context"

	"github.com/anz-bank/sysl-go/core"

	"hello/internal/gen/pkg/servers/hello"
	"hello/internal/handlers"
)

/* AppConfig is where you define app-level config fields. */
type AppConfig struct {
}

func main() {
	hello.Serve(context.Background(),
		func(ctx context.Context, config AppConfig) (*hello.ServiceInterface, *core.Hooks, error) {
			// Perform one-time setup based on config here.
			return &hello.ServiceInterface{
				GetHelloList: handlers.GetHelloListRequest,
			}, nil, nil
		},
	)
}
