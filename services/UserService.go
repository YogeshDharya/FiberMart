package services 
import (     
  "context"    
_  "time"
 _ "log" 
  "github.com/YogeshDharya/fiberBackend/models"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
//  "go.mongodb.org/mongo-driver/mongo/options"
)   
type User models.User
func CreateUser(client *mongo.Client,dbName ,collectionName string,user *User,ansChan chan <- error)  {
  go func(){
    contxt := context.Background()
    Collection := client.Database(dbName).Collection(collectionName)
  _,err := Collection.InsertOne(contxt, User{})//models.User ? not user paramater above ?   
  if err!= nil {  
      ansChan <- err  
  }  
  }()
  } 

func GetUserById(client *mongo.Client,dbName,collectionName ,id string,ansChan chan<- *User) {
  go func(){ 
  contxt := context.Background()
  ObjectId,erro := primitive.ObjectIDFromHex(id)
  if erro != nil {
    ansChan <- nil
    return 
    }
      userCollection := client.Database(dbName).Collection(collectionName) 
      var filter = bson.M{"_id":ObjectId};//DOUBT ??
    var usr User
      err := userCollection.FindOne(contxt,filter).Decode(&usr)
      if err!= nil {
        ansChan <- nil
        return 
      }
  ansChan <- &usr
  }()
}       
