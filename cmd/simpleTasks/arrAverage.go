//Элементы массива, которые меньше среднего арифметического
//https://gospodaretsva.com/avrg-less.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [10]int{}
	rand.Seed(time.Now().UnixNano())
	var add int
	for i := range arr {
		arr[i] = rand.Intn(50) //diapazon ot 10 do 30
		add += arr[i]
	}
	average := add / len(arr)
	fmt.Println(arr)
	fmt.Println("Suma: ", add)
	fmt.Println("Average of array: ", average)

	for j := range arr {
		if arr[j] < average {
			fmt.Println(arr[j])
		}
	}
}
