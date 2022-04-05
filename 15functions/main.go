package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")

	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, _ := proAdder(2, 3, 4, 6)

	fmt.Println("Pro Result is: ", proResult)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro Result function"
}

func greeter() {
	fmt.Println("Namastey from goLang")
}

func adder(valone int, valtwo int) int {
	return valone + valtwo

}
