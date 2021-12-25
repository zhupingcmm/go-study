package main

import (
	"fmt"
	_ "github.com/go-study/src/examples/module1/init/a"
	_ "github.com/go-study/src/examples/module1/init/b"
)

func init () {
	fmt.Println("init from main")
}

func main () {
	fmt.Print("main function")
}