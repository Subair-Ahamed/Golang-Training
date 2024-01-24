package main

import (
	"fmt"
	"encoding/json"
)

func main(){

    type Human struct{
		Name string `json:"name"` //prints name in the json format
		Age int `json:"age"`
		Job string `json:"job"`
	}

	inp1 := Human{"Prakash",23,"Engineer"}
	data1, err := json.Marshal(inp1) //encoding into json format(syntax : json.Marshal(v interface{}) []byte, error) --> returning byte array and error.

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(data1)) //since it returns a byte format, we have to type convert it to string.

	inp2 := []byte(`{"Name":"Ahamed","Age":22,"Job":"AI"}`) //backtick must
	data2 := &Human{} 

	err1 := json.Unmarshal(inp2,data2) // decoding - taking byte format as argument (syntax: json.Unmarshal(data []byte, v interface{}) error) --> returning error.
	err1 := json.Unmarshal(inp2,data2) // decoding - taking byte format as argument (syntax: json.Unmarshal(data []byte, v interface{}) error)
	

	if err1 != nil{
		fmt.Println(err1)
	}

	fmt.Println(data2.Name)
	fmt.Println(data2.Job)
	fmt.Println(data2.Age)

}