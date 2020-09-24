package test

/*import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/51103220/microbon/text"
	//"github.com/51103220/poc.golang/service"
	"net/http"
	"reflect"
	"testing"
)*/

/*
func TestGbonRequestDeserializer(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"CardNumber": "41111111111",
		"ExpiryDate": "123",
	})

	request, _ := http.NewRequest("", "", bytes.NewBuffer(requestBody))
	val := text.DeserializeRequest(nil, request, reflect.TypeOf(service.CardAuthorization{}))

	if val.IsNil() {
		t.Errorf("expected deserializable")
	}
}

func TestGbonRequestDeserializerWithValue(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"CardNumber": "41111111111",
		"ExpiryDate": "123",
	})
	request, _ := http.NewRequest("", "", bytes.NewBuffer(requestBody))

	val := text.DeserializeRequest(nil, request, reflect.TypeOf(service.CardAuthorization{}))

	method := val.MethodByName("Process")
	values := method.Call([]reflect.Value{})

	result := values[1].Elem().Interface().(string)
	fmt.Println(val)
	fmt.Println(result)
	if result != "now ok, 41111111111" {
		t.Errorf("expected field value")
	}
}

func BenchmarkDeserializer(b *testing.B) {
	requestBody, _ := json.Marshal(map[string]string{
		"CardNumber": "41111111111",
		"ExpiryDate": "123",
	})
	request, _ := http.NewRequest("", "", bytes.NewBuffer(requestBody))

	for i := 0; i < b.N; i++ {
		text.DeserializeRequest(nil, request, reflect.TypeOf(service.CardAuthorization{}))
	}
}
*/
