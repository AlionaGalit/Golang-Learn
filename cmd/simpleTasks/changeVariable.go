package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter 2 numbers:")
	fmt.Scanln(&x, &y)

	/*	temp: = x
		x = y
		y = temp*/
	x += y
	y = x - y
	x -= y
	fmt.Printf("x= %d y=%d\n", x, y)
}
