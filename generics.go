package main

import "fmt"

func printValue[T any](value T) {  //T is a type of any type, this is a generic function that can accept any type of value
	fmt.Println(value)
}

func printSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}
func main() {
	printValue(100)
	printValue("Hello")
	printValue(true)
}