package main
import (
	"fmt"
)

func main () {
	x := doOperation(2, increase);
	y :=doOperation(2, decrease);
	fmt.Println(x)
	fmt.Println(y)
}

func increase (a, b int) int{
	fmt.Println("incerase result is ", a + b);
	return a + b;
}

func decrease (a, b int) int{
	fmt.Println("decrease result is ", a + b);
	return a - b;
}

func doOperation (y int, f func(int, int)int) int{
	return f(y,1)
}