package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/YogeshDharya/fiberBackend/controllers"
	// "github.com/YogeshDharya/fiberBackend/models"
	"github.com/YogeshDharya/fiberBackend/routes"
	"github.com/YogeshDharya/fiberBackend/secrets"
	"github.com/YogeshDharya/fiberBackend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/arsmn/fiber-swagger/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var db *mongo.Database
var MongoClient *mongo.Client
func main() {
	fmt.Println("Fiber Mart Starting!")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get port from environment or set default
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Port:", port)

	// Initialize Fiber app
	app := fiber.New()

	// Apply middleware
	app.Use(compress.New()) // Compress responses
	app.Use(cors.New())     // Enable CORS
	app.Static("/docs","./docs")
	// Initialize MongoDB
	MongoClient = InitializeMongoDB()
	defer MongoClient.Disconnect(context.Background())
	db := MongoClient.Database(secrets.Config.DbName)
	userService := services.NewUserService(db,secrets.Config.UserCollection)
	productService := services.NewProductService(db,secrets.Config.ProductCollection)

	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	// Set up routes
	app.Get("/swagger/*", swagger.New(swagger.Config{
    URL:         "/docs/swagger.json",
    DeepLinking: false,
//    URL:         "/swagger/doc.json",
	}))
	// app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
    //     return c.SendFile("./docs/swagger.json")
    // })
	routes.V1Routes(app,userController,productController)

	// Handle non-v1 routes with 404
	app.Use(func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		if len(path) < 3 || path[:3] != "/v1" {
			return ctx.Status(fiber.StatusNotFound).SendString("NOT FOUND!")
		}
		return ctx.Next()
	})

	// Start server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting Fiber Mart: %v", err)
	}
}

func InitializeMongoDB() *mongo.Client{
	uri := secrets.Config.MongoDBURI
	fmt.Printf("Attempting to connect with uri: '%s' \n", uri)
	if uri == ""{
		log.Fatal("Empty DB URI !")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err:= mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error connecting to DB : %v",err)
	}
	MongoClient = client
	err = client.Ping(ctx,nil)
	if err != nil{
		log.Fatalf("Error pinging DB : %v", err)
	}
	fmt.Println("Fiber Mart Started Successfully")

	// defaultDb := secrets.Config.DbName
	// userCollection := secrets.Config.UserCollection
	// productCollection := secrets.Config.ProductCollection
	// cartCollection := secrets.Config.CartCollection
	// Initialize models
	// models.InitializeCartModel(uri, defaultDb, cartCollection)
	// models.InitializeUserModel(uri, defaultDb, userCollection)
	// models.InitializeProductModel(uri, defaultDb, productCollection)
	return MongoClient
}
