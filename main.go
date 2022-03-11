package main

import (
	"context"
	"fmt"
	"mental-health-api/pkg/chat"
	configs "mental-health-api/pkg/config"
	"mental-health-api/pkg/database"
	"mental-health-api/pkg/firebase"
	"mental-health-api/pkg/routes"

	//"mental-health-api/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically

	_ "mental-health-api/docs"
)

// @title Mental Health Api
// @version 1.0
// @description Swagger for Mental Health Api App
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	instance := database.GetMongoInstance()
	defer instance.Client.Disconnect(context.Background())

	config := configs.FiberConfig()
	app := fiber.New(config)

	routes.UserRouter(app)
	routes.PostRouter(app)
	routes.ChatRoutes(app)
	routes.SwaggerRoutes(app)
	routes.UserFeelRouter(app)

	go chat.HubConn.Run()
	firebase.ConnectFirebase()
	firebase.ConnectFirestore()

	fmt.Println("Connected to MongoDB!")

	// Start server (with or without graceful shutdown).
	/*if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}*/
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Listen(port)
}
