package routes

import (
	"github.com/YogeshDharya/fiberBackend/controllers"
	"github.com/gofiber/fiber/v2"
) 
func V1Routes(app *fiber.App,userController *controllers.UserController, productController *controllers.ProductController) {
	v1 := app.Group("/v1")
	V1UserRoutes(v1,userController)
	// userRouter := V1UserRoutes(app,userController)
//	v1.Mount("/user", userRouter)
	V1ProductRoutes(v1,productController)
	//BUG
	//v1.Mount("/products", productRouter)
}

//PENDING :rt should have had been *fiber.Router but works fine when kept as *fiber.Group
// *fiber.Router is of the interface type which include both Group and App .Both fibr.Group and fiber.App satisfy the fiber.Router interface meaning u can pass either of em to a () which expects fiber.Router . TRUE ?