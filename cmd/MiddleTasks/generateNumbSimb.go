//Случайные числа и символы
//https:gospodaretsva.com/random.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//random int
	var max, min int
	fmt.Println("enter min diapozon:")
	fmt.Scanln(&min)
	fmt.Println("enter max diapozon:")
	fmt.Scanln(&max)
	arr := [10]int{}
	for i := range arr {
		arr[i] = min + rand.Intn(max-min)

	}
	fmt.Println("Random int: ", arr)

	//random Float64
	var maxf, minf float64
	fmt.Println("enter min diapozon:")
	fmt.Scanln(&minf)
	fmt.Println("enter max diapozon:")
	fmt.Scanln(&maxf)
	arrf := [10]float64{}
	for i := range arrf {
		arrf[i] = minf + rand.Float64()*(maxf-minf)

	}
	fmt.Printf("Random float64: %.2f \n", arrf)

	//random simbols
	var maxc, minc string
	fmt.Println("enter min caracter diapozon:")
	fmt.Scanln(&minc)
	fmt.Println("enter max caracter diapozon:")
	fmt.Scanln(&maxc)
	min := int(minc[0])
	max := int(maxc[0])

	arrc := [10]int{}
	for i := range arrc {
		arrc[i] = min + rand.Intn(max-min+1)

	}
	for i := range arrc {
		fmt.Printf("%#U\n", arrc[i])
	}
}
