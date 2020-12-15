// Code generated by sysl DO NOT EDIT.
package hello

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Hello
type Service interface {
	GetHelloList(ctx context.Context, req *GetHelloListRequest) (*Message, error)
}

// Client for Hello API
type Client struct {
	Client  *http.Client
	URL     string
	Headers map[string][]string
}

// NewClient for Hello
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL, nil}
}

// GetHelloList ...
func (s *Client) GetHelloList(ctx context.Context, req *GetHelloListRequest) (*Message, error) {
	required := []string{}
	var okResponse Message
	u, err := url.Parse(fmt.Sprintf("%s/hello", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Hello <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkMessageResponse, ok := result.Response.(*Message)
	if ok {
		valErr := validator.Validate(OkMessageResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkMessageResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}