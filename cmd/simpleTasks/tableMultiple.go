//Вывести на экран таблицу умножения
package main

import "fmt"

func multipl(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println() //perehod na novuiu stroku posle pervogo reada
		for j := 1; j <= num; j++ {
			//	multi := i * j
			fmt.Printf("%3d  ", i*j) //multi

		}
	}
}
func main() {

	multipl(9)
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------------------------")

	multipl(20)
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------------------------")

	multipl(5)
}
