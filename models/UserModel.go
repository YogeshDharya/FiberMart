package models

import (
	"context"
	"fmt"
  "os"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//defines structures of the data present as documents within the UserCollection
type User struct{ //struct fields are mapped to fields of MongoDB docx fields  
  Id     primitive.ObjectID `bson:"_id,omitempty"`
  Name   string `bson:"name"`
  Email  string `bson:"email"`
  Password string `bson:"password"`
  Address string `bson:"address"`
  Money string `bson:"money"`
}
var UserCollection *mongo.Collection
func InitializeUserModel(uri,dbName,collectionName string) {
    client,err := mongo.Connect(context.Background(),options.Client().ApplyURI(uri))
  if err!= nil {
    panic(err)
  }
  err = client.Ping(context.Background(),nil)
  if err != nil {
    panic(err)
  }
  database := client.Database(dbName)
  UserCollection := database.Collection(collectionName)//User collection left ?
  fmt.Fprintf(os.Stderr,"%v is Alive\n",UserCollection)//PENDING 
  //index createion , validation or other setups 
}
func Newuser(name,email,password,address,money string) *User{//constructor even requried ?? 
  return &User{Name: name , Email: email , Password: password , Address: address , Money: money}
}















