//Перевести байты в килобайты или наоборот
package main

import "fmt"

func main() {
	var num float32
	var size string
	fmt.Println("Enter number: ")
	fmt.Scanln(&num)
	fmt.Println("Choose change b - byte, or k - kilobyte")
	fmt.Scanln(&size)
	switch size {
	case "b":
		fmt.Println(num * 1024)
	case "k":
		fmt.Printf("%.4f\n", num/1024)
	}
}
