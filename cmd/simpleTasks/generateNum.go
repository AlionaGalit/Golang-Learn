//Вероятность четных случайных чисел
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generatorDiv2 := 0

	rand.Seed(time.Now().UnixNano()) //generiruet cisla v raznom poreadke
	for i := 0; i <= 99; i++ {
		generator := rand.Intn(100)
		fmt.Printf("%v ", generator)
		if i%2 == 0 {
			generatorDiv2 += 1
		}
	}
	fmt.Println("Generator: ", generatorDiv2)
	var procent float32
	procent = float32(generatorDiv2) / 100 * 100 //generator privodim k float32, tak kak procent tipa float32
	fmt.Printf("procent: %.2f%% ", procent)      //%%- eto znak procenta
}
