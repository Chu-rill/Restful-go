package main

import (
	"log"

	"github.com/Chu-rill/Restful-go/database"
	"github.com/Chu-rill/Restful-go/routes"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/api",welcome)
	//User endpoints
	app.Post("/api/users",routes.CreateUser)
	app.Get("/api/users",routes.GetUsers)
	app.Get("/api/users/:id",routes.GetUser)
	app.Put("/api/users/:id",routes.UpdateUser)
	app.Delete("/api/users/:id",routes.DeleteUser)
	//Product endpoint
	app.Post("/api/products",routes.CreateProduct)
}

func welcome(c *fiber.Ctx)error{
	return c.SendString("Welcom to my awesome api")
}


func main(){
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

