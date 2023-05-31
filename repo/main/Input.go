package main

import "fmt"

func main() {
	fmt.Print("Enter a Number : ");
	var input int
	fmt.Scanln(&input)

	if(input % 2 == 0){
		fmt.Printf("Number is Even.");
	}else{
		fmt.Println("Number is Odd.")
	}
	
}