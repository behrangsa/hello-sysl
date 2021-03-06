// Code generated by sysl DO NOT EDIT.
package hello

import (
	"context"
	"time"
)

// DefaultHelloImpl ...
type DefaultHelloImpl struct {
}

// NewDefaultHelloImpl for Hello
func NewDefaultHelloImpl() *DefaultHelloImpl {
	return &DefaultHelloImpl{}
}

// ServiceInterface for Hello
type ServiceInterface struct {
	GetHelloList func(ctx context.Context, req *GetHelloListRequest) (*Message, error)
}

// DownstreamConfig for Hello
type DownstreamConfig struct {
	ContextTimeout time.Duration `mapstructure:"contextTimeout" yaml:"contextTimeout"`
}
