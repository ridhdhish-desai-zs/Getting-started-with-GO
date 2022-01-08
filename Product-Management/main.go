package pm

import "fmt"

func main() {
	var products []Product

	product := AddProduct(1, "New Product")
	products = append(products, product)

	fmt.Println("Products: ", products)
}
