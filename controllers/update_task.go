package controllers

import (
	"VERBA-Test/database"
	"VERBA-Test/models"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateTask(ctx *fiber.Ctx) error {
	log.Println("Received task update request")

	// get task id
	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		log.Println("Error: Invalid ID in URL parameters")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Задача не найдена",
		})
	}

	//parse request
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		log.Println("Error: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неправильный формат данных",
		})
	}

	// query existing task from DB
	var task models.Task
	queryResult := database.DB.First(&task, id)
	if queryResult.Error == gorm.ErrRecordNotFound {
		log.Println("Task not found")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Задача не найдена",
		})
	} else if queryResult.Error != nil {
		log.Println("Error: Failed query data from DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	// create new task and fill it with updated data
	updatedTask := task

	value, exists := data["due_date"]
	if exists {
		// validate due_date field
		_, err := time.Parse(time.RFC3339, data["due_date"])
		if err != nil {
			log.Println("Error: Field [due_date] is invalid")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Неправильный формат данных.",
			})
		}
		updatedTask.Due_date = value
	}
	value, exists = data["title"]
	if exists {
		updatedTask.Title = value
	}
	value, exists = data["description"]
	if exists {
		updatedTask.Description = value
	}

	// update record
	queryResult = database.DB.Model(task).Omit("created_at").Updates(updatedTask)
	if queryResult.Error != nil {
		log.Println("Error: Failed to update data in DB")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Проблема на сервере",
		})
	}

	// data updated successfully
	log.Println("Operation completed successfully")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":          updatedTask.ID,
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"due_date":    updatedTask.Due_date,
		"created_at":  updatedTask.Created_at,
		"updated_at":  updatedTask.Updated_at,
	})
}
