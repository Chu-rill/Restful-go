package routes

import (
	"github.com/Chu-rill/Restful-go/database"
	"github.com/Chu-rill/Restful-go/models"
	"github.com/gofiber/fiber/v2"
)


type User struct{
	//this is not the model User,see this as the serializer
	ID uint `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func createResponseUser(userModel models.User)User{
return User{ID: userModel.ID,FirstName: userModel.FirstName,LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx)error{
	var user models.User
	
	if err := c.BodyParser(&user);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	database.Databse.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(fiber.StatusCreated).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx)error{

}