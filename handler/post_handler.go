package handler

import (
	models "mental-health-api/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(ctx *fiber.Ctx) error {
	var post models.Post

	if err := ctx.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := post.Create(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	return ctx.Status(fiber.StatusCreated).JSON(post)
}

func GetPost(ctx *fiber.Ctx) error {
	var post models.Post
	post_id := ctx.Params("postid")

	result, err := post.GetOne(post_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func GetPosts(ctx *fiber.Ctx) error {
	var post models.Post

	results, err := post.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(results)
}

func DeletePost(ctx *fiber.Ctx) error {
	/*var post models.Post
	post_id := ctx.Params("postid")

	err := post.DeleteOne(post_id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(200)*/
	var post models.Post
	post_id := ctx.Params("postid")

	post.BaseModel.Deleted = true
	post.BaseModel.DeletedAt = time.Now().Unix()

	err := post.DeleteOne(post_id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(200)
}

/*func DeletePosts(ctx *fiber.Ctx) error {
	var post models.Post

	err := post.DeleteAll()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(200)
}*/

func UpdatePost(ctx *fiber.Ctx) error {
	var post models.Post
	post_id := ctx.Params("postid")

	if err := ctx.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := post.UpdateOne(post_id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(200)
}
