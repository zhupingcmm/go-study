package main

import (
	"fmt"
	"sync"
)

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	
}


func loopFunc() {
	lock := sync.Mutex{}
	for i :=0; i<3; i++ {
		go func(i int)   {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)	
	}
}