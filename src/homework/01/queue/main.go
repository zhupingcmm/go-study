package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	queue := make(chan int, 10)

	defer close(queue)
	go func() {
		c := time.NewTicker(1*time.Second)
		for temp := range c.C {
			select {
			default:
				fmt.Printf("%v received message: %d\n", temp, <-queue)
			}
		}
	}()

	ticker := time.NewTicker(1*time.Second)
	for next := range ticker.C {
		num := rand.Intn(100)
		queue <- num
		fmt.Printf("%v send message: %d \n", next, num)
	}


}