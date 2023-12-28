package routes
import (
  "github.com/gofiber/fiber/v2"
)//PENDING :rt should have had been *fiber.Router but works fine when kept as *fiber.Group
//*fiber.Router is of the interface type which include both Group and App .Both fibr.Group and fiber.App satisfy the fiber.Router interface meaning u can pass either of em to a () which expects fiber.Router . TRUE ?
func V1Routes(app *fiber.App) {   
  v1 := app.Group("/v1")
  userRouter := V1UserRoutes(app)
  v1.Mount("/user",userRouter) 
  productRouter := V1ProductRoutes(app)
  //BUG 
  v1.Mount("/products",productRouter)
}  
