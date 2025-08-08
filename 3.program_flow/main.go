package main

import (
	"fmt"
)

func main() {

	/*
		price, inStock := 100, true

		if price > 80 {
			fmt.Println("too expensive")
		}
		_ = inStock

		//

		if price <= 100 && inStock == true {
			fmt.Println("buy it")
		}

		//

		if price < 100 {
			fmt.Println("its cheap")
		} else if price == 100 {
			fmt.Println("on the edge")
		} else {
			fmt.Println("its expensive")
		}

		//  voting example

		age := 5

		if age > 0 && age < 18 {
			fmt.Printf("you cannot vote , please return in %d years \n ", 18-age)
		} else {
			fmt.Println("you can vote")
		}

	*/

	//  ================.  input via command line argument ======

	//  via os package
	// fmt.Println("os args ", os.Args)

	//  os .Args is a slice , so we can also access particular argument with index also
	//  like , sso 0th argument is path of os args , and 1st index is 1st argument

	// fmt.Println("first cli argument is ", os.Args[1])

	//  ============= LOOPS

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//  go doesnot have while loop , but we can make for loop act as while , in below
	j := 10
	for j > 0 {
		fmt.Println(j)
		j--
	}

}
