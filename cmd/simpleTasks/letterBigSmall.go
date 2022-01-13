//Количество строчных и прописных букв в строке
package main

import "fmt"

func main() {
	var line string
	fmt.Println("Enter line: ")
	fmt.Scanln(&line)

	big := 0
	small := 0
	for i := range line {
		if line[i] >= 'a' && line[i] <= 'z' {
			small += 1
		} else if line[i] >= 'A' && line[i] <= 'Z' {
			big += 1
		}
	}
	fmt.Println("big: ", big)
	fmt.Println("small: ", small)
}
