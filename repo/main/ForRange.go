package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Sum : ", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index : ", i)
		}
	}
	kvs := map[string]string{"1": "Dileep", "2": "Babji", "3": "Sujatha"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\v", k, v)
	}
}
