//Поменять столбцы матрицы местами
//https://gospodaretsva.com/change-columns.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [5][5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		for j := range arr {
			arr[i][j] = rand.Intn(100)
		}
		fmt.Printf("%3d \n", arr[i])
	}
	var column1, column2 int
	fmt.Println("Enter number of colomn would you like change: ")
	fmt.Scanln(&column1)
	fmt.Scanln(&column2)

	for i := range arr {

		for j := range arr {
			j = column1
			temp := arr[i][j]
			arr[i][j] = arr[i][column2]
			arr[i][column2] = temp

		}
		fmt.Printf("%3d \n", arr[i])
	}

}
