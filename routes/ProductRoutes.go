
package routes
import (
  "github.com/gofiber/fiber/v2"
  "github.com/YogeshDharya/fiberBackend/controllers"
)
func V1ProductRoutes(router fiber.Router,productController *controllers.ProductController){
  productRouter := router.Group("/products")
  productRouter.Get("/",productController.GetAllProductHandler)
}
// app.Get("/",productController.GetAllProductHandler)