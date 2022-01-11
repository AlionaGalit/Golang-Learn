//Количество двузначных чисел в матрице
//https://gospodaretsva.com/double-digit.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [10][10]int{}
	rand.Seed(time.Now().UnixNano())

	for i := range arr { //row
		for j := range arr { //column
			arr[i][j] = rand.Intn(1000)
			fmt.Printf("%3d  ", arr[i][j])
		}
		fmt.Println()
	}

	count := 0
	for i := range arr {
		for j := range arr {
			if arr[i][j] > 9 && arr[i][j] < 100 {
				count += 1
			}
		}
	}

	fmt.Println("Double numbers: ", count)

}
