package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRouter(a *fiber.App) {
	router := a.Group("/post")
	router.Get("/", handler.GetPosts)
	router.Get("/:postid?", handler.GetPost)
	router.Post("/create", handler.CreatePost)
	router.Put("/update/:postid?", handler.UpdatePost)
	router.Put("/delete/:postid?", handler.DeletePost)
}
