package goRoutines

import (
	"fmt"
)

func Numbers(ch chan<- int) {
	for i := 2; i < 100; i++ {
		ch <- i
	}
}

func PrimeNumbers(in <-chan int, out chan<- int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

func EX() {
	numbers := make(chan int)
	go Numbers(numbers)

	for i := 0; i < 10; i++ {
		prime := <-numbers
		fmt.Println(prime)
		newChannel := make(chan int)
		go PrimeNumbers(numbers, newChannel, prime)
		numbers = newChannel
	}
	fmt.Println("calling the parllel sum function :")
	Parllel()

}
