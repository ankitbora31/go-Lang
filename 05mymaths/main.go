package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("the sum is: ", mynumberOne+int(mynumberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from cypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
