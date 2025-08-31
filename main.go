package main

import (
	//"context"
	"fmt"
	//"log"
	_ "time"
	"github.com/joho/godotenv"
	"github.com/YogeshDharya/fiberBackend/models"
	"github.com/YogeshDharya/fiberBackend/routes"
	"github.com/YogeshDharya/fiberBackend/secrets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
  "github.com/gofiber/fiber/v2/middleware/cors"
//  "github.com/gofiber/fiber/v2/middleware/bodyparser"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongodb+srv://yogeshdharya:Dante$1998@cluster0.izilr8v.mongodb.net/?retryWrites=true&w=majority
var db *mongo.Database

func main() {
	fmt.Println("LetzStore MyKart !")
	app := fiber.New()
	if err:= godotenv.Load() ; err != nil {
		panic(err)
	}
	port := os.Getenv("DB_PORT")// STORE_PORT -> DB_PORT 
	fmt.Println("Port Before godotenv.Write()",port)
	godotenv.Write(map[string][string]{"DB_PORT":key , 3000})
	newPort := os.Getenv("DB_PORT");
	fmt.Println("Port After writing into .env file",newPort)
   app.Use(bodyparser.New()) //for the post request bodies parsing 
	app.Use(compress.New())   //compresses response
  InitializeMongoDB()

  app.Use(cors.New())       //for cors   

  routes.V1Routes(app)
  app.Use(func(ctx *fiber.Ctx) error {``  
      if ctx.Path()[:3] != "/v1" {//3's exclusive     
        return ctx.Status(fiber.StatusNotFound).SendString("NOT FOUND !") 
      }  
     return ctx.Next()
      })
	errs := app.Listen(":",newPort)     //TODO 
	if errs != nil {//log vs panic vs Errorf ?? PENDING      
 		fmt.Errorf("Error : %v", errs)     
	}    
}      
func InitializeMongoDB()  {  
  uri := secrets.Config.MongoDBURI
  defaultDb := secrets.Config.DbName
  userCollection := secrets.Config.UserCollection
  productCollection := secrets.Config.ProductCollection
 cartCollection := secrets.Config.CartCollection
  models.InitializeCartModel(uri,defaultDb,cartCollection) 
  models.InitializeUserModel(uri,defaultDb,userCollection) 
  models.InitializeProductModel(uri,defaultDb,productCollection)
}
