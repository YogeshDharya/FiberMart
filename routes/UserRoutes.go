package routes 
import (  
  "github.com/gofiber/fiber/v2"
"github.com/YogeshDharya/fiberBackend/controllers"
) 
func V1UserRoutes(app *fiber.App) *fiber.App{ 
  app.Get("/:uuid",controllers.ByIdHandler)
  app.Post("/",controllers.CreateUserHandler)
  return app
}