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

	// validate request fields
	_, exists := data["title"]
	if !exists {
		log.Println("Error: Field [title] in request body is absent.")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных.",
		})
	}
	_, exists = data["description"]
	if !exists {
		log.Println("Error: Field [description] in request body is absent.")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных.",
		})
	}
	_, exists = data["due_date"]
	if !exists {
		log.Println("Error: Field [due_date] in request body is absent.")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных.",
		})
	}
	_, err := time.Parse(time.RFC3339, data["due_date"])
	if err != nil {
		log.Println("Error: Field [due_date] is invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных.",
		})
	}
	// можно еще добавить проверку на попытку создать задачу с истекшим сроком
	// но на данный момент это фича

	newTask := models.Task{
		Title:       data["title"],
		Description: data["description"],
		Due_date:    data["due_date"],
		Created_at:  time.Now().Format(time.RFC3339),
		Updated_at:  time.Now().Format(time.RFC3339),
	}

	// save data to DB
	if err := database.DB.Create(&newTask).Error; err != nil {
		log.Println("Error: Failed to write data to DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	// data saved successfully
	log.Println("Operation completed successfully")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":          newTask.ID,
		"title":       newTask.Title,
		"description": newTask.Description,
		"due_date":    newTask.Due_date,
		"created_at":  newTask.Created_at,
		"updated_at":  newTask.Updated_at,
	})

}
