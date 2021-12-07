package main

import "fmt"

func main() {
	var a, b, c, m int32
	fmt.Println("Enter 3 numbers:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	m = a
	if m < b {
		m = b
	} else if m < c {
		m = c
	}
	fmt.Println("The max number :", m)

}
