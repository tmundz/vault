package handler

import (
  "github.com/gofiber/fiber/v2"
)

func testHandler(c *fiber.Ctx) error {
  return c.SendString("api says hi")
}
