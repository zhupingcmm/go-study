package a

import (
	"fmt"
	_ "github.com/go-study/src/examples/module1/init/b"
)

func init () {
	fmt.Println("init form a")
}