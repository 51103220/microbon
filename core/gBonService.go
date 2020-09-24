package core

import (
	"net/http"
	"reflect"
)

type GBonServiceFunc func(in *GBonRequest) (error, interface{})

type GBonServiceFilter struct {
	Auth            http.HandlerFunc
	RequestFilters  ByPriority
	ResponseFilters ByPriority
}

type GBonService struct {
	Path        string
	Verb        string
	PayloadType reflect.Type
	Filter      *GBonServiceFilter
}

func (service *GBonService) Authenticate(authFunc http.HandlerFunc) *GBonService {
	service.Filter.Auth = authFunc
	return service
}

func (service *GBonService) RequestFilter(priority int, requestFilterFunc http.HandlerFunc) *GBonService {
	filter := &GBonFilter{
		Priority: priority,
		Handler:  requestFilterFunc,
	}

	service.Filter.RequestFilters.AddFilter(filter)

	return service
}

func (service *GBonService) ResponseFilter(priority int, responseFilterFunc http.HandlerFunc) *GBonService {
	filter := &GBonFilter{
		Priority: priority,
		Handler:  responseFilterFunc,
	}

	service.Filter.ResponseFilters.AddFilter(filter)

	return service
}
