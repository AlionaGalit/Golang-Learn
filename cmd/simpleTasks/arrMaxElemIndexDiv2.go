//Максимальный из элементов массива с четными индексами
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [10]int{}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		arr[i] = rand.Intn(50)
	}
	fmt.Println(arr)

	var max int
	for i := 0; i < len(arr); i += 2 {
		if i == 0 {
			max = arr[i]
			continue //v etom meste idet na proverku cikla i dobavleat sag+2
		}
		if max < arr[i] {
			max = arr[i]
		}

	}
	fmt.Println("Max ciotnii element: ", max)

}
