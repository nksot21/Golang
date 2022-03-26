package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRouter(a *fiber.App) {
	router := a.Group("/post")
	router.Get("/", handler.GetPosts)
	router.Get("/:postid?", handler.GetPost)
	router.Post("/", handler.CreatePost)
	router.Put("/:postid?", handler.UpdatePost)
	router.Delete("/:postid?", handler.DeletePost)
	router.Get("/top5/:emotion?", handler.Get5Posts)
}
