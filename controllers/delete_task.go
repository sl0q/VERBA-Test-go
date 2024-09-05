package controllers

import (
	"VERBA-Test/database"
	"VERBA-Test/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteTask(ctx *fiber.Ctx) error {
	log.Println("Received delete task by ID request")

	// get task id
	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		log.Println("Error: Invalid ID in URL parameters")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Задача не найдена",
		})
	}

	// query tasks from DB
	var task models.Task
	queryResult := database.DB.First(&task, id)
	if queryResult.Error == gorm.ErrRecordNotFound {
		log.Println("Task not found")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Задача не найдена",
		})
	} else if queryResult.Error != nil {
		log.Println("Error: Failed to query data from DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	queryResult = database.DB.Delete(&task)
	if queryResult.Error != nil {
		log.Println("Error: Failed to query data from DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	log.Println("Operation completed successfully")
	return ctx.Status(fiber.StatusNoContent).JSON("")
}
