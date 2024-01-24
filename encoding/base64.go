package main

import (
	"fmt"
	"encoding/base64"
)

func main(){
	text := "Hello! This is base64 encoding!!!"
	enc := base64.StdEncoding.EncodeToString([]byte(text)) //to encode the string into base64
	fmt.Println(enc)

	dec,err := base64.StdEncoding.DecodeString(enc) //decode from base64
	if err != nil{
		fmt.Println("Error decoding",err)
    }
	fmt.Println(string(dec)) 
}