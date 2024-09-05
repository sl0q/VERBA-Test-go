package controllers

import (
	"VERBA-Test/database"
	"VERBA-Test/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ViewTask(ctx *fiber.Ctx) error {
	log.Println("Received view task by ID request")

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

	log.Println("Operation completed successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":          task.ID,
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.Due_date,
		"created_at":  task.Created_at,
		"updated_at":  task.Updated_at,
	})
}
