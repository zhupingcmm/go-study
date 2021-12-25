package main
import (
	"fmt"
)

func main () {
	x := doOperation(2, increase);
	y := doOperation(2, decrease);
	fmt.Println(x)
	fmt.Println(y)
}

func increase (a, b int) int {
	return a+b;
}

func decrease(a,b int) int {
	return a-b;
}

func doOperation (a int, f func(int, int)int) int {
	return f(a, 1)
}