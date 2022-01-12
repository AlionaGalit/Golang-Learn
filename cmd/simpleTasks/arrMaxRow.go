// Дана таблица (матрица) 10*10.
//Найти строку сумма чисел которой максимальная.
//Заполнение матрицы - случайное (положительные и отрицательные числа.)
//Тип чисел любой
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [10][10]int{}
	arrMax := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		for j := range arr {
			arr[i][j] = rand.Intn(100)
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Println()
	}

	for i := range arr {
		sum := 0
		for j := range arr {
			sum += arr[i][j]

		}
		arrMax = append(arrMax, sum)
	}
	fmt.Println(arrMax)
	//fmt.Printf("%3d   \n", arrMax)

	maxElem := 0
	index := 0
	for i := range arrMax {
		if maxElem < arrMax[i] {
			maxElem = arrMax[i]
			index = i // v kakoi stroke max Suma
		}

	}
	fmt.Printf("Sum Max %d in the %d row ", maxElem, index)
}
