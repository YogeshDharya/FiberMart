package controllers

import (
	//"log"
	"github.com/YogeshDharya/fiberBackend/services"
	"github.com/YogeshDharya/fiberBackend/models"
	"github.com/gofiber/fiber/v2"
//	"go.mongodb.org/mongo-driver/mongo"
	//  "github.com/YogeshDharya/fiberBackend/secrets"
)

type UserController struct{
	userService *services.UserService
}

func NewUserController(	userService *services.UserService) *UserController{
	return &UserController{
		userService: userService,
	}
}
// These are the files which directly interact with client APIs 29th Dec 2024 -->                                                                                                                 Dekha Dekha mujhey bhi bunney luga . May be its fine to behave like a child at times to be happy n sad within a second is far more healthy than lasting periods for the same
func (userController *UserController) ByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id Absent"})
	}
	ch := make(chan *models.User)

	go userController.userService.GetUserById(id,ch)
	var user = <-ch
	if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return ctx.JSON(user)
} 
func (userController *UserController) CreateUserHandler(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	ansChan := make(chan error, 1)
	userService := userController.userService
	go userService.CreateUser(&user, ansChan)
	if err := <-ansChan; err != nil {
		return nil
	}

	return nil
}
