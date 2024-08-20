package main

import (
	"fmt"
	"golangfiles/calculation"
)

func main() {
	fmt.Println("hello from go ")

	x, y := 7, 20

	fmt.Println("Performimg Calculation")
	fmt.Println("Calcution is", calculation.Cal(x, y))

	var i int

	fmt.Println("Print even number ")
	for i = 10; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}

}
