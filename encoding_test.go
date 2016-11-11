package encoding

import (
	"testing"
)

type DumyStruct struct {
	i     int
	f     float32
	b     bool
	Slice []int
	Map   map[string]interface{}
	Ptr   *string
}

func TestMarshal(t *testing.T) {
	// test int
	mapTest := make(map[string]interface{})
	mapTest["kint"] = int(10)
	mapTest["ki8"] = int8(8)
	mapTest["kstring"] = "string value"
	ptrSting := "Ptr String"
	dumy := DumyStruct{
		i:     1,
		f:     2.0,
		b:     true,
		Slice: []int{1, 2, 3},
		Map:   mapTest,
		Ptr:   &ptrSting,
	}
	target := "{\"i\": 1,\"f\": 2.000000,\"b\": true,\"Slice\": [1,2,3],\"Map\": {\"kint\":10,\"ki8\":8,\"kstring\":\"string value\"},\"Ptr\": \"Ptr String\"}"
	bytes, err := Marshal(dumy)
	if err != nil || string(bytes) != target {
		t.Error(err)
	}
}
