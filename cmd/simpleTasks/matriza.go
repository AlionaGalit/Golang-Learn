//Дана таблица (матрица) 10*10. Найти строку сумма чисел которой максимальная.
//Заполнение матрицы - случайное (положительные и отрицательные числа.)
//Тип чисел любой
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	matrix := [10][10]int{}
	rand.Seed(time.Now().UnixNano())

	for i := range matrix {
		for j := range matrix {
			matrix[i][j] = rand.Intn(5) //generatia sluceinih cisel do 5
		}
	}

	sumMax := 0
	rowMaxIndex := 0
	for i := range matrix {
		sum := 0
		for j := range matrix {
			sum += matrix[i][j]
		}
		fmt.Printf("%v Row - Suma: %v\n", i, sum)
		if sumMax < sum {
			sumMax = sum
			rowMaxIndex = i
		}
	}
	fmt.Printf("%v Row - Suma max: %v\n", rowMaxIndex, sumMax)

	fmt.Println(matrix)

}
