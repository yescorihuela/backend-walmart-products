package domain

type ProductRepositoryMongo struct{}

func (prm ProductRepositoryMongo) GetAllProducts() ([]Product, error) {
	return nil, nil
}

func (prm ProductRepositoryMongo) GetProduct(criteria string) (*Product, error) {
	return nil, nil
}

func NewProductRepositoryMongo(db *interface{}) ProductRepositoryMongo {
	return ProductRepositoryMongo{}
}
