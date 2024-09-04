package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	return ctx.SendString("<VERBA-group test")
}
