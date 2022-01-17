package domain

import (
	"context"
	"strconv"
	"time"

	"github.com/yescorihuela/walmart-products/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryMongo struct {
	client *mongo.Client
}

func (prm ProductRepositoryMongo) GetAllProducts() ([]Product, *errs.AppError) {
	var products []Product
	collection := prm.client.Database("promotions").Collection("products")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &products); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if len(products) == 0 {
		return nil, errs.NewNotFoundError("Elements not found")
	}

	return products, nil
}

func (prm ProductRepositoryMongo) GetProductsByCriteria(criteria string) ([]Product, *errs.AppError) {
	var products []Product
	collection := prm.client.Database("promotions").Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var filter bson.M

	id, errId := strconv.Atoi(criteria)
	if errId == nil {
		filter = bson.M{"id": id}
	} else {
		filter = bson.M{
			"$or": bson.A{
				bson.M{"brand": bson.M{"$regex": criteria, "$options": "i"}},
				bson.M{"description": bson.M{"$regex": criteria, "$options": "i"}},
			},
		}
	}

	productsByCriteria, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	defer productsByCriteria.Close(ctx)
	if err = productsByCriteria.All(ctx, &products); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if len(products) == 0 {
		return nil, errs.NewNotFoundError("Elements not found")
	}

	return products, nil
}

func NewProductRepositoryMongo(db *mongo.Client) ProductRepositoryMongo {
	return ProductRepositoryMongo{client: db}
}
