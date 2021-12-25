package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		fmt.Println("hello from thread")
		ch <- 0
	}()
	i := <-ch
	fmt.Println(i)
}