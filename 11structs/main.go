package main

import "fmt"

func main() {
	fmt.Println("Structs in goLang")

	// no inheritance, no super or parent in goLang

	ankit := User{"Ankit", "ankit31@go.dev", true, 22}
	fmt.Println(ankit)

	fmt.Printf("Ankit Details are: %+v\n", ankit)
	fmt.Printf("Name is %v and Email is %v\n", ankit.Name, ankit.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
