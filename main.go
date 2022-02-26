package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	configs "mental-health-api/pkg/config"
	"mental-health-api/pkg/database"
	"mental-health-api/pkg/routes"
	"mental-health-api/pkg/utils"
	"os"
)

func main() {
	instance := database.GetMongoInstance()
	defer instance.Client.Disconnect(context.Background())

	config := configs.FiberConfig()
	app := fiber.New(config)

	routes.UserRouter(app)

	fmt.Println("Connected to MongoDB!")

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
