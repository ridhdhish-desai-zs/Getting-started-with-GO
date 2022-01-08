package pm

type Product struct {
	id   int
	name string
}

var products = []Product{}

func GetProducts() []Product {
	return products
}

func AddProduct(id int, name string) Product {
	var newProduct Product

	newProduct.id = id
	newProduct.name = name

	products = append(products, newProduct)

	return newProduct
}

func DeleteProduct(id int) Product {
	for k, v := range products {
		if v.id == id {
			return products[k]
		}
	}

	return Product{}
}
