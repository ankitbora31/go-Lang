package main

import "fmt"

func main() {
	fmt.Println("Maps in goLang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("After deletion: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
}
