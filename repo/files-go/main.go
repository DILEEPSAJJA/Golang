package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Learning About Files.")
	content := "I am Currently working on files topic in golang. "

	file , err := os.Create("./example.txt")

	if err != nil {
		panic(err)
	}

	length , err := io.WriteString(file, content)

	if err != nil{
		panic(err)
	}
	fmt.Println("The Length is : ", length)
	defer file.Close()
}
