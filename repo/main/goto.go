package main

import "fmt"

func main(){

	var input int
Loop:
		fmt.Println("You Are Not eligible for vote.")
		fmt.Print("Enter Your Age : ")
		fmt.Scanln(&input)
		if(input <=18){
			goto Loop
		}else{
			fmt.Println("You can Vote.")
		}
}