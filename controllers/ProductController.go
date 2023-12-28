package controllers
import(
  "github.com/gofiber/fiber/v2"
  "github.com/YogeshDharya/fiberBackend/services"
)
func GetAllProductHandler(ctx *fiber.Ctx) error{
  products,err := services.GetAllProducts()
  if err != nil {
    return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed To Fetch Products"})
  }
  return ctx.Status(fiber.StatusOK).JSON(products)
}