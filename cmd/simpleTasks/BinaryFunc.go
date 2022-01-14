//Функция перевода десятичного числа в двоичное
//https://gospodaretsva.com/binary.html
package main

import (
	"fmt"
)

func Binary(decNum int) string {
	arr := []int{}
	for decNum > 0 {
		arr = append(arr, decNum%2) //результат записываем в слайс

		decNum = decNum / 2

	}
	var s string
	for i := len(arr) - 1; i >= 0; i-- { //обходим слайс с конца
		s += fmt.Sprintf("%d", arr[i]) //и переводим в string исп-я Sprintf

	}
	return s
}

func main() {
	for {
		var decNum int

		fmt.Println("Enter decemal number: ")
		fmt.Scanln(&decNum)
		if decNum == 0 { //если вводим 0 выход из прогораммы
			break
		}
		fmt.Println(Binary(decNum))

	}
}
