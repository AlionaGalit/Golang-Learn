//Вывести на экран таблицу умножения
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		fmt.Println() //perehod na novuiu stroku posle pervogo reada
		for j := 1; j <= 9; j++ {
			//	multi := i * j
			fmt.Printf("%3d  ", i*j) //multi

		}
	}
}
