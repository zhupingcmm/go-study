package main

import (
	"fmt"
)

func main() {
	fmt.Print("abc")
	result := add(1,2);
	fmt.Println(result)
	err := fmt.Errorf("this is a error")
	fmt.Println(err)
}

func add (a, b int) int {
	return a + b;
}