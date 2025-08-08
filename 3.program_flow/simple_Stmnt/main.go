package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println("ERROR : ", err)
	} else {
		fmt.Println(i)
	}

	//  lets write the upper code in simple statement or short statement
	if i, err := strconv.Atoi("34"); err == nil {
		println(i)
	} else {
		println(err)
	}

}
