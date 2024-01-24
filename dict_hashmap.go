package main

import "fmt"

func main() {

	stationary := map[string]int{ //syntax : map[key_type]value_type{key-value pairs}
		"pencil": 2,
		"pen":  4,
		"marker": 3,
		"scale":  4,
	}
	stationary["chart"] = 5 //to add new element(key-value pair) to the map

	delete(stationary, "scale") //to delete a pair
	fmt.Println(stationary)

    things,ok := stationary["chart"] //ok checks whether true or false(chart is present or not)...
	if ok{
		fmt.Println(things,"Charts are there")
	} else {
		fmt.Println("No  charts")
	}

	for x, y := range stationary { //iterating through the map(just like foreach)
		fmt.Printf("%v : %v\n", x, y)
	}

	//to create a map using make
	bike := make(map[string]string) 
	bike["brand"] = "yamaha"
	bike["model"] = "mt15"
	bike["year"] = "2019"
	fmt.Println(bike)

	bike["model"]+="new version" //assignment operation
	fmt.Println(bike)

	//Example to get input from user:
	var input string
	fmt.Print("Enter the input(pencil,pen,marker,scale,chart):")
	fmt.Scanf("%s", &input)
	display := stationary[input]
	fmt.Printf("No. of things are %d\n", display)

}