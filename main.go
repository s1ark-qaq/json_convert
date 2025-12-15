package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	marshalStruct()

	marshalNorStruct()
	
	unmarshalToStruct()

	unmarshalToMap()
}

func marshalStruct() {
	var u = &Student{
		Name:   "s1ark",
		Age:    20,
		IsGood: true,
	}
	resp, _ := json.Marshal(u)
	fmt.Println(string(resp))
}

func marshalNorStruct() {
	var u = struct {
		Name string
		Age  int
	}{
		Name: "s1ark",
		Age:  20,
	}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}

func unmarshalToStruct() {
	jsonStr := `{"Name":"张三","Age":21,"IsGood":true}`
	var u Student
	json.Unmarshal([]byte(jsonStr), &u)
	fmt.Println(u)
}

func unmarshalToMap() {
	jsonStr := `{"Name":"s1ark","Age":23,"Interests":{"Sports":["Run","Jump"]}}`

	data := make(map[string]any, 0)
	json.Unmarshal([]byte(jsonStr), &data)

	fmt.Println(data["Name"].(string))
	fmt.Println(data["Interests"].(map[string]any)["Sports"].([]any))
	fmt.Println(data)
}
