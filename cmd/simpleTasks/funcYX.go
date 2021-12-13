//Вычислить значение функции y=f(x)
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Println("Enter x: ")
	fmt.Scanln(&x)
	if x > 0 {
		y = 2*x - 10
	} else if x == 0 {
		y = 0
	} else if x < 0 {
		y = 2*math.Abs(x) - 1
	}
	fmt.Printf("Result: %v", y)
}
