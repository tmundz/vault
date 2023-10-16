package main

import (
  //"fmt"
  //"log"
  "github.com/gofiber/fiber/v2"
  "github.com/tmundz/Vault/routes"
)


func main() {
  app := fiber.New()
  app.Get("/", func(c * fiber.Ctx) error {
    return c.SendString("Hello")
  })
  app.Listen(":8080")
}
