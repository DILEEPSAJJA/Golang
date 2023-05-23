package main

import (
	"fmt"
)

func main(){
	fmt.Println("map -  Introduction.")

	lang := make(map[string]string)

	lang["JS"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["PY"] = "python"

	fmt.Println("List of languages : " , lang)

	for key, value := range lang{
		fmt.Printf("For key %v , For Value %v\n",key,value)
	}

}