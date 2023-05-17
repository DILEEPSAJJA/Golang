package main

import ("fmt"
		"os"
		"bufio"
	)

func main(){
	fmt.Println("Welcome to go lang tutorial")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	fmt.Println("Rating : ",input)
	
}