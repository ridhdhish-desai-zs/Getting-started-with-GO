package pm

import "errors"

type Product struct {
	id   int
	name string
}

var products = []Product{}

var conn string = "This is connection string"

func GetProducts() ([]Product, error) {
	if conn == "" {
		connectionError := errors.New("CONNECTION HAS BEEN LOST")
		return nil, connectionError
	}

	return products, nil
}

func AddProduct(id int, name string) (Product, error) {

	var newProduct Product

	if conn == "" {
		connectionError := errors.New("CONNECTION HAS BEEN LOST")
		return newProduct, connectionError
	}

	newProduct.id = id
	newProduct.name = name

	products = append(products, newProduct)

	return newProduct, nil
}

func DeleteProduct(id int) (Product, error) {
	if conn == "" {
		connectionError := errors.New("CONNECTION HAS BEEN LOST")
		return Product{}, connectionError
	}

	for k, v := range products {
		if v.id == id {
			return products[k], nil
		}
	}

	return Product{}, nil
}
