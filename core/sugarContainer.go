package core
import "reflect"

type SugarContainer struct {
	FunctionName        string
	Version string
	PayloadType reflect.Type
}