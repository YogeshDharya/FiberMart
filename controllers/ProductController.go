package controllers
import(
  "github.com/gofiber/fiber/v2"
  "github.com/YogeshDharya/fiberBackend/services"
)

type ProductController struct{
  productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController{
  return &ProductController{
    productService: productService,
  }
}

func (productController *ProductController) GetAllProductHandler(ctx *fiber.Ctx) error{
  products,err := services.GetAllProducts()
  // products,err := services.GetAllProducts()
  if err != nil {
    return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed To Fetch Products"})
  }
  return ctx.Status(fiber.StatusOK).JSON(products)
}