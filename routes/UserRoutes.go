package routes 
import (  
  "github.com/gofiber/fiber/v2"
  "github.com/YogeshDharya/fiberBackend/controllers"
) 
// func V1UserRoutes(app *fiber.App,userController *controllers.UserController) *fiber.App{ 
func V1UserRoutes(router fiber.Router,userController *controllers.UserController){ 
  userRouter := router.Group("/user")
  userRouter.Get("/:uuid",userController.ByIdHandler)
  userRouter.Post("/",userController.ByIdHandler)
//  return app
}