package routes

import (
	"github.com/Chu-rill/Restful-go/database"
	"github.com/Chu-rill/Restful-go/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product)Product{
	return Product{
		ID: productModel.ID,
		Name: productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx)error{
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Databse.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(201).JSON(responseProduct)
}