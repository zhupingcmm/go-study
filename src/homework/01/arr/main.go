package main

import "fmt"

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for k, v := range arr {
		if v == "stupid" {
			arr[k] = "smart"
		}
		if v == "weak" {
			arr[k] = "strong"
		}
	}

	fmt.Println(arr)
}