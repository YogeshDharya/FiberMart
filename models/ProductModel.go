package models

import (
	"context"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Product struct {
  Id     primitive.ObjectID `bson:"_id,omitempty"`
  Name   string `bson:"name"`
  Category string `bson:"category"`
  Cost int `bson:"cost"`
  Rating int `bson:"rating"`
  image string `bson:"image"`
}
var ProductCollection *mongo.Collection 
func InitializeProductModel(uri,dbName,collectionName string) { 
client , err := mongo.Connect(context.Background(),options.Client().ApplyURI(uri))
  if err!= nil {
    panic(err)  
  }
  err = client.Ping(context.Background(),nil)
  if err != nil {
    panic(err)
  }
  database := client.Database(dbName)
  ProductCollection := database.Collection(collectionName)
  fmt.Fprintf(os.Stderr,"%v is Alive ! \n",ProductCollection)
}
func NewProduct(name,category,image string, cost , rating int) *Product{
  return &Product{
    Name : name,
    Category : category,
    Cost : cost ,
    Rating : rating,
    image : image,
  }
}
