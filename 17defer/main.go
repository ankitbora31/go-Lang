package main

import "fmt"

func main() {
	defer fmt.Println("First")
	defer fmt.Println("second")
	defer fmt.Println("Third")
	fmt.Println("fourth")
	fmt.Println("five")

	myDefer()
}

func myDefer() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
