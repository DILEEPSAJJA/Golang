package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f float64 = 9.4
	var str1 string = "101"
	var str2 string = "100"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newint, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newint)

	newfloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newfloat)

}
