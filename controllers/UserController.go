package controllers

import (  
  //"log"
  "github.com/gofiber/fiber/v2"  
  "github.com/YogeshDharya/fiberBackend/services" 
  "github.com/YogeshDharya/fiberBackend/secrets/config"
) 
// These are the files which directly interact with client APIs 29th Dec 2024 -->                                                                                                                 Dekha Dekha mujhey bhi bunney luga . May be its fine to behave like a child at times to be happy n sad within a second is far more healthy than lasting periods for the same 
func ByIdHandler(ctx *fiber.Ctx) error {
  id := ctx.Params("id") 
  if id == ""{
     return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id Absent"})
    } 
  ch := make(chan *services.User) 
  
  go services.GetUserById(secrets.Config.,ch)  
    var user = <- ch 
    if user == nil {     
    return  ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})  
   }   
  return ctx.JSON(user)   
}   
func CreateUserHandler(ctx *fiber.Ctx) error{ 
   var user services.User 
  if err:= ctx.BodyParser(&user); err != nil{
    return ctx.Status(fiber.StatusBadRequest).JSON(err)   
  }  
  go services.CreateUser(ctx.Body(),user)   
  err:= <-  
  return ctx.JSON(user); 
} 
  