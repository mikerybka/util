package main

import (
	"fmt"

	"encoding/json/jsontext"

	"github.com/mikerybka/util"
)

func main() {
	data := `{"one":[1,23,23142],"two":"hi","three":false,"four":123,"five":{"yo":"ho"}}`
	var v jsontext.Value
	if err := v.UnmarshalJSON([]byte(data)); err != nil {
		panic(err)
	}
	err := util.WriteJSONFS("d/b", v)
	if err != nil {
		fmt.Println(err)
	}
}
