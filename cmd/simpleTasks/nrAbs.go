//Номер минимального по модулю элемента массива
package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []float64{-10, -2, 8, -6, 2, 4, 2, -2}
	numMin := math.Abs(slice[0])

	var p int //index
	for i := 1; i < len(slice); i++ {
		if numMin > math.Abs(slice[i]) {
			numMin = math.Abs(slice[i])
			p = i
		}
	}
	fmt.Printf("Min number: %v\nIndex: %v \n", numMin, p) //numMin soderjit znacenie, p- index

	for j := 0; j < len(slice); j++ {
		if math.Abs(slice[j]) == numMin {
			fmt.Printf("Index: %v\n", j)

		}
	}
	//	fmt.Printf("Min number: %v %v\nIndex: %v %v", numMin,  p, b) //numMin soderjit znacenie, p- index

}
