package services

import (
	"context"
	"github.com/YogeshDharya/fiberBackend/models"
	"go.mongodb.org/mongo-driver/bson"
)   
var newProduct models.Product
func GetAllProducts() ([]models.Product,error){
  ctx := context.Background() 
  iterator,err := models.ProductCollection.Find(ctx,bson.M{})
  if err != nil {
    return nil,err
  }
  var products []models.Product 
  if err:= iterator.All(ctx,&products); err != nil {
    return nil,err
  }
  return products ,nil
}
func GeProductById(id string){
  
  }