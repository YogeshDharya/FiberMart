package services

import (
	"context"

	"github.com/YogeshDharya/fiberBackend/models"
	"go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

//var newProduct models.Product

type ProductService struct {
	collection *mongo.Collection
}
func NewProductService(db *mongo.Database, collection string)  *ProductService{
  return &ProductService{
    collection: db.Collection(collection),
  }
}
func GetAllProducts() ([]models.Product, error) {
	ctx := context.Background()
	iterator, err := models.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var products []models.Product
	if err := iterator.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func GeProductById(id string) {

}
