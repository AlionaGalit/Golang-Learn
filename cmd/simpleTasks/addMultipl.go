//Сумма и произведение цифр числа
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number: ")
	fmt.Scanln(&num)
	add := 0
	multipl := 1

	for num > 0 {
		add += num % 10
		multipl *= num % 10
		num = num / 10
	}

	fmt.Printf("Suma: %v\nMultipl: %v\n", add, multipl)
}
