package registry

import (
	"fmt"
	"github.com/51103220/microbon/core"
	"reflect"
)

var functionHolder = make(map[string]*core.SugarContainer)

func RegisterFunction(dto core.Executor) {
	var function = &core.SugarContainer{
		FunctionName: dto.FunctionName(),
		Version:      dto.GetApiVersion(),
		PayloadType:  reflect.TypeOf(dto),
	}

	functionHolder[fmt.Sprintf("%s:%s", dto.FunctionName(), dto.GetApiVersion())] = function
}

func GetFunctions() map[string]*core.SugarContainer {
	return functionHolder
}
