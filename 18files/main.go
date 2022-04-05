package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GOLang")

	content := "this need to go in a file - files"

	// creation is os operation
	file, err := os.Create("./myGofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./myGofile.txt")
}

// for reading ioutil
func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	// fmt.Println("Text data in file is: \n", databyte)
	fmt.Println("Text data in file is: \n", string(databyte))

}

// handle error
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
