package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m := [10][10]int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {

			m[i][j] = 100 - rand.Intn(200) //diapazon -100 do 100
		}
	}
	for _, row := range m {
		for _, elem := range row { //cifra zanimaet 2 pozitii
			fmt.Printf("%3d ", elem)
		}
		fmt.Println()
	}
	fmt.Println("Diagonali: ")
	for i := 0; i <= 9; i++ {
		if m[i][i] > 0 {
			fmt.Printf("%v ", m[i][i])
		}
	}

}
