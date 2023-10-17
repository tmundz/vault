package main

import (
  //"fmt"
  //"log"
  "os"
  "github.com/gofiber/fiber/v2"
  "github.com/tmundz/Vault/api/routes"
)

func getEnv(key, defaulValue string) string {
  if value, exists := os.LookupEnv(key); exists {
    return value
  }
  return defaulValue
}

func main() {
  server := getEnv("SERVER_ADDRESS", "8080")
  app := fiber.New()
  app.Listen(server)
}
