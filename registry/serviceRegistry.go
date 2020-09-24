package registry

import (
	"github.com/51103220/microbon/core"
	"github.com/51103220/microbon/helper"
	"reflect"
)

var serviceHolder = make(map[string]*core.GBonService)

func Post(dto core.GBonRequest) *core.GBonService {
	return initialize(dto, "post")
}

func Get(dto core.GBonRequest) *core.GBonService {
	return initialize(dto, "get")
}

func GetServices() map[string]*core.GBonService {
	return serviceHolder
}

func initialize(dto core.GBonRequest, verb string) *core.GBonService {
	path := helper.NormalizePath(dto.RegisterAs())
	var service = &core.GBonService{
		Path:        path,
		Verb:        verb,
		PayloadType: reflect.TypeOf(dto),
		Filter:      &core.GBonServiceFilter{},
	}

	serviceHolder[path] = service

	return service
}
