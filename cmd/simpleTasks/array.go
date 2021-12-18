//Заполнение и вывод массивов
//https://gospodaretsva.com/filling.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr1 := [10]int{}
	rand.Seed(time.Now().UnixNano()) //vvodim sliceainim pereborom
	for i := range arr1 {
		arr1[i] = rand.Intn(100)
	}
	fmt.Printf("%3d \n", arr1)

	arr2 := [10]int{}
	var num int
	for i := range arr2 { //vvodim s kliviaturi
		fmt.Print("Enter number: ")
		fmt.Scan(&num)
		arr2[i] = num
	}
	fmt.Printf("%3d \n", arr2)

	arr3 := [10]int{}
	for i := range arr3 { //summa 1 i 2 massiva zapisivaetsea v tretii
		arr3[i] = arr1[i] + arr2[i]
	}
	fmt.Printf("%3d \n", arr3)

}
