package text

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func DeserializeRequest(w http.ResponseWriter, r *http.Request, t reflect.Type) reflect.Value {
	ptr := reflect.New(t)
	s := ptr.Elem().Addr().Interface()
	json.NewDecoder(r.Body).Decode(&s)
	return ptr
}
