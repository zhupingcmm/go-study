package main
import (
	"fmt"
)

func main () {
	var arr [5]string = [5]string{"I", "am", "stupid", "and", "weak"};
	for k, v:= range arr {
		if (v=="stupid") {
			arr[k] = "smart";
		}
		if (v == "weak") {
			arr[k] = "strong";
		}	
	}
	fmt.Print(arr)
}