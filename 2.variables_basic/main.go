package main

import "fmt"

func main() {
	//  1st type of var declaration :
	// var age int = 10

	//  2nd type of var declaration
	var age = 10

	// 3rd type of var declaration
	name := "Dany"
	fmt.Println("Name , age is ", name, age)

	//

	//  multiple var delarations
	car, cost := "Audi", 500000

	fmt.Println("car , cost : ", car, cost)

	car, year := "BMW", 2002
	fmt.Println(car, year)

	//  other way
	var i, j int
	i, j = 1, 2
	_, _ = i, j // to mute the warning

}
