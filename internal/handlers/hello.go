package handlers

import (
	"context"
	"hello/internal/gen/pkg/servers/hello"

	"github.com/anz-bank/sysl-go/common"
)

// GetHelloListRequest returns a hello message
func GetHelloListRequest(ctx context.Context, getHelloListRequest *hello.GetHelloListRequest) (*hello.Message, error) {

	// Set response content type to JSON
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}

	// return the result
	return &hello.Message{
		Body: "Hello, World!",
	}, nil
}
