//Вывести ряд чисел в диапазоне с шагом
package main

import "fmt"

func main() {
	var minNum, maxNum, stepNum int
	fmt.Println("Enter min number: ")
	fmt.Scanln(&minNum)
	fmt.Println("Enter max number: ")
	fmt.Scanln(&maxNum)
	fmt.Println("Enter step number: ")
	fmt.Scanln(&stepNum)

	for i := minNum; i <= maxNum; i += stepNum {

		fmt.Printf("%v ", i)
	}

}
