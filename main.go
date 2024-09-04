package main

import (
	"fmt"
	"os"

	"VERBA-Test/database"
	"VERBA-Test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: VERBA-Test [databaseCredentialsFilePath.json]")
		return
	}
	databaseCredentialsFilePath := os.Args[1]

	_, err := database.ConnectDB(databaseCredentialsFilePath)

	if err != nil {
		panic("Failed to connect to Database")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:     "POST, GET, PUT, DELETE",
		AllowHeaders:     "Content-Type",
		AllowCredentials: false,
	}))

	routes.SetUpRoutes(app)

	err = app.Listen(":8000")

	if err != nil {
		panic("Failed to start server")
	}
}
