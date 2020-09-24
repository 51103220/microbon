package test

import (
	"context"
	"fmt"
	"github.com/51103220/microbon/registry"
	"reflect"
	"testing"
)

type somethingImplementsGbonRequest struct {
	cardNumber string
	expiryDate string
}

func (c somethingImplementsGbonRequest) RegisterAs() string {
	return "/dummy/service/registry"
}

func (c somethingImplementsGbonRequest) GetApiVersion() string {
	return "1009"
}

func (c somethingImplementsGbonRequest) Process(ctx context.Context) (error, interface{}) {
	return nil, fmt.Sprintf("now ok, %s", c.cardNumber)
}

func TestPostServiceRegister(t *testing.T) {
	service := registry.Post(somethingImplementsGbonRequest{})

	if service == nil {
		t.Errorf("expected non-nil")
	}

	if service.Verb != "post" {
		t.Errorf("expected post")
	}
}

func TestGetServiceRegister(t *testing.T) {
	service := registry.Get(somethingImplementsGbonRequest{})

	if service == nil {
		t.Errorf("expected non-nil")
	}

	if service.Verb != "get" {
		t.Errorf("expected get")
	}
}

func TestServiceRegisterType(t *testing.T) {
	service := registry.Post(&somethingImplementsGbonRequest{})

	if service == nil {
		t.Errorf("expected non-nil")
	}

	if service.PayloadType != reflect.TypeOf(&somethingImplementsGbonRequest{}) {
		t.Errorf("expected valid type")
	}
}

func BenchmarkPostServiceRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		registry.Post(&somethingImplementsGbonRequest{})
	}
}
