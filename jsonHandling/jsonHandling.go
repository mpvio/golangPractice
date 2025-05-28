package jsonHandling

import (
	"encoding/json"
	"fmt"
)

type sample struct {
	// field names must start with capital letters to be exported!
	// otherwise they are ignored!
	Name string
	Num  int `json:"num"` //this makes it so the json version will still be saved in lowercase
}

func MarshalDataToJson() {
	s := sample{Name: "hello", Num: 15}

	encoding, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Json:", string(encoding))

	sSli := []sample{{"hi", 2}, {"3", 3}}

	sliEnc, err := json.Marshal(sSli)
	if err != nil {
		panic(err)
	}
	fmt.Println("even collection of structs can be marshalled:", string(sliEnc))
}

func UnmarshalJsonToData() {
	//data as bytes
	//fields here are case insensitive
	data := []byte(`{
	"name": "hello",
	"num": 10}`)

	//need to define variable/ "holder" ahead of time
	var obj sample

	//decoding function takes a reference to 'holder'
	err := json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println("success:", obj.Name)

	// can also unmarshal arrays
	dataArr := []byte(`[
	{"name": "hi",
	"num": 10},
	{"name": "bye",
	"num": 20}]`)

	var objs []sample

	err2 := json.Unmarshal(dataArr, &objs)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("success:", objs[0].Name, objs[1].Name)
}
