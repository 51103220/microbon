package registry

import (
	"fmt"
	"github.com/51103220/microbon/core"
	"reflect"
)

var functionHolder = make(map[string]*core.SugarContainer)

func RegisterFunction(dto core.Executor) {
	var functionContainer = &core.SugarContainer{
		FunctionName: dto.FunctionName(),
		Version:      dto.GetApiVersion(),
		PayloadType:  reflect.TypeOf(dto),
	}

	functionHolder[createFuncKey(dto.FunctionName(), dto.GetApiVersion())] = functionContainer
}

func GetFunctionContainer(functionName string, version string) *core.SugarContainer {
	if version == "" {
		version = "1"
	}

	functionContainer, ok := functionHolder[createFuncKey(functionName, version)]

	if ok {
		return functionContainer
	}

	return nil
}

func createFuncKey(functionName string, version string) string {
	return fmt.Sprintf("%s:%s", functionName, version)
}
