package b

import (
	"fmt"
)

func init() {
	fmt.Println("init form b")
	
}

func main() {
	a := 
	a,b = passVal(1,2)
	fmt.Println(a,b)
}


func passVal (a,b int) (x,y int) {
	x=a;
	y=b;
	return x,y;
}