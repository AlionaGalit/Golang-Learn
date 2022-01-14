//Функция заполнения массива случайными числами
//https://gospodaretsva.com/array.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Array(start int, finish int) [10]int {
	arr := [10]int{}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		arr[i] = start + rand.Intn(finish-start)

	}
	return arr

}
func main() {

	var min, max int
	fmt.Println("Enter min numberof diapazon:")
	fmt.Scanln(&min)
	fmt.Println("Enter max numberof diapazon:")
	fmt.Scanln(&max)
	a := Array(min, max)
	fmt.Println(a)
}
