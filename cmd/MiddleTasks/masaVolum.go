//Известны данные о массе (в кг) и объеме (в см3) 20 тел,
//изготовленных  из  различных  материалов.
//Определить  макси-мальную плотность материала
package main

import (
	"fmt"
	"math/rand"
)

type Density struct {
	Masa    float32
	Volum   float32
	Densiti float32
}

func main() {
	//	rand.Seed(time.Now().UnixNano())
	products := []Density{} //Density это тип (как int, float)
	for i := 0; i <= 20; i++ {
		p := Density{ //в р хранится вся структура
			Masa:  rand.Float32() * 100, //заполнение случайными значениями
			Volum: rand.Float32() * 100,
		}
		products = append(products, p)

	}
	fmt.Printf("%.2f\n", products)
	///////////////////////////
	var maxMasaProduct Density
	for i := 0; i <= 20; i++ { //нахожу максимальную массу

		if maxMasaProduct.Masa < products[i].Masa {
			maxMasaProduct = products[i]
		}
	}
	fmt.Println()
	fmt.Println("Max masa ", maxMasaProduct)
	///////////////////////////
	//нахожу плотность материала

	for i := 0; i <= 20; i++ {
		products[i].Densiti = products[i].Masa / products[i].Volum
	}
	fmt.Printf("%.2f", products)

	var maxDensiti Density
	for i := 0; i <= 20; i++ {
		if maxDensiti.Densiti < products[i].Densiti {
			maxDensiti = products[i]
		}
	}
	fmt.Println()
	fmt.Println("Density max: ", maxDensiti)
}
