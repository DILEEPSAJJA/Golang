package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("file in golang")
	content := "My name is Dileep Syudying in B.Tech 2nd Year IT-A."

	file, err:= os.Create("./exam.txt")

	if err != nil{
		panic(err)
	}

	length, err:= io.WriteString(file, content)
	if err != nil{
		panic(err)
	}
	fmt.Println("The LEngth is : ", length)
	file.Close()
}