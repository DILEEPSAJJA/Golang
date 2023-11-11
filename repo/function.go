package main

import "fmt"

type Employee struct{
	fname string
	lname string
}

func (emp Employee) fullname(){
	fmt.Println(emp.fname + "  " + emp.lname )
}
func main(){
	e1 := Employee {"Dileep","Sajja"}
	e1.fullname()
}
