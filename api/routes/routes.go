package api

import (
  "github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
  app.Get("/"), func(c *fiber.Ctx) error {
    return c.SendString("hello")
  })
}

