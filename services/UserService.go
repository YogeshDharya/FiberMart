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
type  UserServiceI interface {
    CreateUwser(cli mongo.Client,dbName string,collectionName string,usr *models.User,chanValue chan<-error )
    GetUserById(client *mongo.Client,dbName,collectionName ,id string,ansChan chan<- *models.User)
}
type UserService struct{
  collection *mongo.Collection
}
func NewUserService(db *mongo.Database, collection string) *UserService{
  return &UserService{
    collection: db.Collection(collection),
  }
}
func (service *UserService) CreateUser(user *models.User,ansChan chan <- error)  {
  go func(){
    contxt := context.Background()
    collection := service.collection
  _,err := collection.InsertOne(contxt, models.User{})//models.User ? not user paramater above ?   
  if err!= nil {  
      ansChan <- err  
  }  
  }()
  } 

func (service *UserService) GetUserById( id string,ansChan chan<- *models.User) {
  go func(){ 
  ctx := context.Background()
  ObjectId,error := primitive.ObjectIDFromHex(id)
  if error != nil {
    ansChan <- nil
    return 
    }
      userCollection := service.collection
      var filter = bson.M{"_id":ObjectId};//DOUBT ??
    var usr models.User
      err := userCollection.FindOne(ctx,filter).Decode(&usr)
      if err!= nil {
        ansChan <- nil
        return 
      }
  ansChan <- &usr
  }()
}       
