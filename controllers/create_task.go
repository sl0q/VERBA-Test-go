package controllers

import (
	"VERBA-Test/database"
	"VERBA-Test/models"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(ctx *fiber.Ctx) error {
	log.Println("Received tokens giveaway request")

	//parse request
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных",
		})
	}

	newTask := models.Task{
		Title:       data["title"],
		Description: data["description"],
		Due_date:    data["due_date"],
		Created_at:  time.Now().Format(time.RFC3339),
		Updated_at:  time.Now().Format(time.RFC3339),
	}

	// save data to DB
	if err := database.DB.Create(&newTask).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	// Выдача токенов прошла успешно
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":          newTask.ID, // CHANGE
		"title":       newTask.Title,
		"description": newTask.Description,
		"due_date":    newTask.Due_date,
		"created_at":  newTask.Created_at,
		"updated_at":  newTask.Updated_at,
	})

}
