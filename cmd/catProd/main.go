//Есть данные по категориям.категория содержит имя, описание,  список товаров.
//каждый товар содержит название, кол-во,цену, список магазинов где он продаётся
//каждый магазин содержит название, город и адрес.
//1. создать структуру к-ая описывает инф-ию одной категории
//2. заполнение списка категорий
package main

import "fmt"

type Category struct {
	Name        string
	Description string
	Products    []Product
}

type Product struct {
	Name     string
	Quantity int
	Price    float32
	Shops    []Shop
}

type Shop struct {
	Name   string
	Town   string
	Adress string
}

func main() {
	// input information
	categories := WriteCategories() //use function WriteCategories

	fmt.Println(categories)
}

func WriteCategories() []Category {
	categories := []Category{}
	//start write data
	var cat int
	fmt.Println("Enter number of categories: ")
	fmt.Scanln(&cat)
	for i := 1; i <= cat; i++ {
		var catTemp Category
		fmt.Printf("Enter data of category %d\n", i)
		fmt.Print("Name: ")
		fmt.Scanf("%s\n", &catTemp.Name)
		fmt.Print("Description: ")
		fmt.Scanf("%s\n", &catTemp.Description)

		//   add products
		catTemp.Products = AddProducts() //use function AddProducts
		////////////////////////////
		categories = append(categories, catTemp)
		fmt.Println("-------------------------------------")
	}
	//---end write data
	return categories
}

func AddProducts() []Product {
	var products []Product
	var prodNumber int
	fmt.Println("How much products will be added: ")
	fmt.Scanln(&prodNumber)
	for j := 1; j <= prodNumber; j++ {
		var prodTemp Product
		fmt.Printf("Enter data of product %d\n ", j)
		fmt.Print("Name: ")
		fmt.Scanf("%s\n", &prodTemp.Name)
		fmt.Print("Quantity: ")
		fmt.Scanf("%d\n", &prodTemp.Quantity)
		fmt.Print("Price: ")
		fmt.Scanf("%f\n", &prodTemp.Price)

		//added shops
		prodTemp.Shops = AddShops() //use function AddShops
		/////////////////////
		products = append(products, prodTemp)
	}
	return products
}

func AddShops() []Shop {
	var shops []Shop
	var shopNumber int
	fmt.Print("Enter number of shops: ")
	fmt.Scanf("%d\n", &shopNumber)
	for k := 1; k <= shopNumber; k++ {
		var shopTemp Shop
		fmt.Printf("Enter data shop %d\n", k)
		fmt.Print("Name: ")
		fmt.Scanf("%s\n", &shopTemp.Name)
		fmt.Print("Town: ")
		fmt.Scanf("%s\n", &shopTemp.Town)
		fmt.Print("Adress: ")
		fmt.Scanf("%s\n", &shopTemp.Adress)
		shops = append(shops, shopTemp)
	}
	return shops

}
