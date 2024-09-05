package controllers

import (
	"VERBA-Test/database"
	"VERBA-Test/models"

	"log"

	"github.com/gofiber/fiber/v2"
)

func ListTasks(ctx *fiber.Ctx) error {
	log.Println("Received tokens giveaway request")

	// query tasks from DB
	var tasks []models.Task
	queryResult := database.DB.Find(&tasks)
	if queryResult.Error != nil {
		log.Println("Error: Failed to query data from DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	// composing responce
	var responce []fiber.Map
	for _, task := range tasks {
		responce = append(responce, fiber.Map{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.Due_date,
			"created_at":  task.Created_at,
			"updated_at":  task.Updated_at,
		})
	}

	log.Println("Operation completed successfully")
	return ctx.Status(fiber.StatusCreated).JSON(responce)
}
