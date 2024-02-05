package pkg2

import (
    "fmt"
)

type SomeStruct struct {
    Field string
}

func (s *SomeStruct) SomeMethod() {
    fmt.Println("This is a method of SomeStruct.")
}
