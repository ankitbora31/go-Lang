package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in goLang")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "banana"

	fmt.Println("FruitList ", fruitList)
	fmt.Println("FruitList size ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list ", vegList)
	fmt.Println("Veg list size is: ", len(vegList))

}
