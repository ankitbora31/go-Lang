package main

import "fmt"

const LoginToken string = "afsdafkjslf"

func main() {
	var username string = "ankit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.2343242
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	// stores 0 when nothing assigned
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "ankitbora31.in"
	fmt.Println(website)

	// no var style
	// walrus operator (not allowed outside function)
	numberOfUser := 400000
	fmt.Println(numberOfUser)

	fmt.Println("login token: ", LoginToken)
}
