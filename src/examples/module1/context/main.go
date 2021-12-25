package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background();
	ctx := context.WithValue(baseCtx, "a", "b");
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	timeoutCtx, cancel:= context.WithTimeout(baseCtx, 5*time.Second)

	defer cancel()

	go func(c context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C{
			select {
			case <-c.Done():
				fmt.Println("child process interrupt...")
			default:
				fmt.Println("enter default")
			}
			
		}
	}(timeoutCtx)

	select {
		case <-timeoutCtx.Done():
			// time.Sleep(1 * time.Second)
			fmt.Println("main process exit!")
	}
}