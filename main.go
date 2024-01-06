package main

import (
	//"context"
	"fmt"
	//"log"
	_ "time"

	"github.com/YogeshDharya/fiberBackend/models"
	"github.com/YogeshDharya/fiberBackend/routes"
	"github.com/YogeshDharya/fiberBackend/secrets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/bodyparser"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongodb+srv://yogeshdharya:Dante$1998@cluster0.izilr8v.mongodb.net/?retryWrites=true&w=majority
var db *mongo.Database

func main() {
	fmt.Println("LetzStore MyKart !")
	app := fiber.New()
   app.Use(bodyparser.New()) //for the post request bodies parsing 
	app.Use(compress.New())   //compresses response
  InitializeMongoDB()

  app.Use(cors.New())       //for cors   

  routes.V1Routes(app)
  app.Use(func(ctx *fiber.Ctx) error {  
      if ctx.Path()[:3] != "/v1" {//3's exclusive     
        return ctx.Status(fiber.StatusNotFound).SendString("NOT FOUND !") 
      }  
     return ctx.Next()
      })
	errs := app.Listen(":3000")     //TODO 
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
//TODO
o  for managing schema i Skeema ii golang-migrate 
o guide to connect go lang to plantescale - https://planetscale.com/docs/tutorials/connect-go-gorm-app 
//package main 

// import ( 
//     "database/sql" 
//     "log" 
//     "os" 

//     "github.com/joho/godotenv" 
//      _ "github.com/go-sql-driver/mysql" 
// ) 
 
// func main() { 
//     // Load connection string from .env file 
//     err := godotenv.Load() 
//     if err != nil { 
//         log.Fatal("failed to load env", err) 
//     } 
 
//     // Open a connection to PlanetScale 
//     db, err := sql.Open("mysql", os.Getenv("DSN")) 
//     if err != nil { 
//         log.Fatalf("failed to connect: %v", err) 
//     } 

//     rows, err := db.Query("SHOW TABLES")
//     if err != nil {
//         log.Fatalf("failed to query: %v", err)
//     }
//     defer rows.Close()

//     var tableName string
//     for rows.Next() {
//         if err := rows.Scan(&tableName); err != nil {
//             log.Fatalf("failed to scan row: %v", err)
//         }
//         log.Println(tableName)
//     }

//     defer db.Close()
// }
