package main

import (
	"fmt"
	"strings" //string packages
)

func main() {

	str1 := "Datasirpi"
	str2 := "Golang"

	/*---String Methods---*/

	fmt.Println(str1 + str2) //String Concatenation

	fmt.Println(strings.Contains(str1, "a")) //if a is in str1 or not

	fmt.Println(strings.Count(str1, "i")) //gives the number of times

	fmt.Println(strings.HasPrefix(str1, "D"))
	fmt.Println(strings.HasSuffix(str2, "g"))

	fmt.Println(strings.Index(str1, "t")) //gives the position of character

	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str2))

	fmt.Println(len(str1))

	fmt.Println(strings.Split(str1, ""))

	fmt.Println(strings.Replace(str2, "G", "M", 1))  //4 arguments is valid (stringname, "character", "replacer", once)
	fmt.Println(strings.Replace(str1, "i", "a", -1)) //-1 indicates similar characters should be replaced

	fmt.Println(strings.Repeat(str2, 2))

	/*---Iterating through the string---*/

	for i := 0; i < len(str1); i++ {

		fmt.Printf("The index of %c is %d\n", str1[i], i)

	}

}
