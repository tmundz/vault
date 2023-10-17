package routes

import (
  "github.com/gofiber/fiber/v2"
  "github.com/tmundz/Vault/api/handler"
)


func SetupRoutes(app *fiber.App) {
  app.Get("/api/test", handler.testHandler)
}

