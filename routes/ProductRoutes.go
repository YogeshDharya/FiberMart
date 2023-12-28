
package routes
import (
  "github.com/gofiber/fiber/v2"
  "github.com/YogeshDharya/fiberBackend/controllers"
)
func V1ProductRoutes(app *fiber.App){
    app.Get("/",controllers.GetAllProductHandler)
}