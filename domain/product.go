package domain

type Product struct {
	Id          uint    `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func NewProduct(id uint, brand, description string, price float32) Product {
	newProduct := Product{
		Id:          id,
		Brand:       brand,
		Description: description,
		Price:       price,
	}
	return newProduct
}
