//Количество положительных, отрицательных и равных нулю элементов массива
//https://gospodaretsva.com/positive-negative.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [20]int{}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		arr[i] = -5 + rand.Intn(4-(-5)) //diapazon ot-5 do 4
	}
	fmt.Printf("%2v\n", arr)
	var neg, poz, zero int
	for j := range arr {
		if arr[j] < 0 {
			neg += 1
		} else if arr[j] > 0 {
			poz += 1
		} else {
			zero += 1
		}
	}
	fmt.Printf("Negative: %v\n", neg)
	fmt.Printf("Pozitive: %v\n", poz)
	fmt.Printf("Zero: %v\n", zero)

}
