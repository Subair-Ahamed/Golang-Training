package main

import "fmt"

func main() {

	wheels := map[string]int{ //syntax : map[key_type]value_type{key-value pairs}
		"bike": 2,
		"car":  4,
		"auto": 3,
		"suv":  4,
	}
	wheels["lorry"] = 6 //to add new element(key-value pair) to the map

	delete(wheels, "suv") //to delete a pair
	fmt.Println(wheels)

	for x, y := range wheels { //iterating through the map
		fmt.Printf("%v : %v\n", x, y)
	}

	cars := make(map[string]string) //to create a map using make()
	cars["brand"] = "audi"
	cars["model"] = "a3"
	cars["year"] = "2018"
	fmt.Println(cars)

	//Example to get input from user:
	var input string
	fmt.Print("Enter the input(car,bike auto,lorry):")
	fmt.Scanf("%s", &input)
	display := wheels[input]
	fmt.Printf("No. of wheels are %d", display)

}
