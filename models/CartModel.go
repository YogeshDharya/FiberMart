package models

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type Cart struct {
  Id primitive.ObjectID `bson:"_id,omitempty"`
  Email string
  paymentOptions string
  CartItems   []Product
}
var CartCollection *mongo.Collection
func InitializeCartModel(uri,dbName,collectionName string){
  client , err := mongo.Connect(context.Background(),options.Client().ApplyURI(uri))
  if err!= nil {
    panic(err)
    }
  err = client.Ping(context.Background(),nil)
  if err != nil {
    panic(err)
    }
  database := client.Database(dbName)
  CartCollection := database.Collection(collectionName)
fmt.Fprintf(os.Stderr,"%v i am alive\n",CartCollection)
} 