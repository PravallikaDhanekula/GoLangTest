package main

import (
	"fmt"
)

func sumRange(start, end int, ch chan<- int) {
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	ch <- total
}

func main() {

	start := 1
	end := 10

	ch := make(chan int, 10)

	mid := (start + end) / 2

	go sumRange(start, mid, ch)
	go sumRange(mid+1, end, ch)

	sum1 := <-ch
	sum2 := <-ch

	totalSum := sum1 + sum2

	fmt.Println("Total Sum:", totalSum)
}
