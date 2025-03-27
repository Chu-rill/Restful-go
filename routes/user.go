package routes

import (
	"errors"

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
	users := []models.User{}

	database.Databse.Db.Find(&users)
	responseUsers := []User{}
	for _,user := range users{
		responseUser := createResponseUser(user)
		responseUsers = append(responseUsers,responseUser)
	}
	return c.Status(fiber.StatusCreated).JSON(responseUsers)

}

func findUser(id int,user *models.User)error{
	 database.Databse.Db.First(&user,"id=?",id)
	if user.ID == 0{
		return errors.New("User not found")
	}	
	return nil

}

func GetUser(c *fiber.Ctx)error{
	id,err := c.ParamsInt("id")

	var user models.User
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	 
	if err := findUser(id,&user);err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}
	responseUser := createResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx)error{
	id,err := c.ParamsInt("id")

	var user models.User
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id,&user);err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}

	type UpdateUser struct{
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser
	if err := c.BodyParser(&updateData);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Databse.Db.Save(&user)

	responseUser := createResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx)error{
	id,err := c.ParamsInt("id")

	var user models.User
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id,&user);err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}

	if err := database.Databse.Db.Delete(&user).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(200).SendString("Succesfully Deleted User")
}