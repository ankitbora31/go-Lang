package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to Slices in goLang")
	// slices are arrays with more features

	var fruitList = []string{"Apple", "Peach", "Orange"}
	fmt.Printf("type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 165
	highScores[3] = 867

	highScores = append(highScores, 555, 665, 231)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println("Sorted? ", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "python", "java", "javascript", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
