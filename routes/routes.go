package routes

import (
	"VERBA-Test/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/tasks", controllers.CreateTask)
	app.Get("/tasks", controllers.ListTasks)
	app.Get("/tasks/{id}", controllers.ViewTask)
	app.Put("/tasks/{id}", controllers.UpdateTask)
	app.Delete("/tasks/{id}", controllers.DeleteTask)
}
