package domain

type ProductRepositoryStub struct {
	products []Product
}

func (p ProductRepositoryStub) GetAllProducts() ([]Product, error) {
	return p.products, nil
}

func (p ProductRepositoryStub) GetProduct(criteria string) (*Product, error) {
	return &Product{Id: 101, Brand: "Amor a Roma", Description: "adda", Price: 29900}, nil
}

func NewRepositoryStub() ProductRepositoryStub {
	products := []Product{
		{Id: 1, Brand: "Hamilton Beach", Description: "Coffee Maker", Price: 49990},
		{Id: 2, Brand: "General Electric", Description: "Microwave oven", Price: 29900},
		{Id: 101, Brand: "Amor a Roma", Description: "adda", Price: 29900},
	}
	return ProductRepositoryStub{products}
}
