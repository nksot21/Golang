package main

import (
	"chatdemo/src/chat"
	"chatdemo/src/firebase"

	"chatdemo/src/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	app := fiber.New()

	//CONNECT DATABASE
	firebase.ConnectFirebase()
	firebase.ConnectFirestore()

	//CREATE CHAT HUB
	go chat.HubConn.Run()
	routes.SetUpChatRoutes(app)
	routes.SetUpUserRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Listen(":8080")
}
