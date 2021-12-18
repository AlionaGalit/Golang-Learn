//Сумма и произведение элементов массива
//https://gospodaretsva.com/sum-mult.html
package main

import "fmt"

func main() {
	var num float32

	arr := [5]float32{}
	for i := range arr {
		fmt.Printf("Enter number %v: ", i)
		fmt.Scanln(&num)
		arr[i] = num
	}
	fmt.Println(arr)
	var add float32 = 0
	var mult float32 = 1
	for j := range arr {
		add += arr[j]
		mult *= arr[j]
	}
	fmt.Println(add)
	fmt.Printf("%.2f", mult)
}
