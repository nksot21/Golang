package handler

import (
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

// Create Post
// @Summary Create Post
// @Description EVENT_emotion = 0
// @Description POST__happy = 1
// @Description POST__sad = 2
// @Description POST__scared = 3
// @Description POST__angry = 4
// @Description POST__worry = 5
// @Description POST__normal = 6
// @Description POST__depression = 7
// @Tags /post
// @Accept json
// @Produce json
// @Param post body models.Post true "Post"
// @Success 200 ""
// @Router /post [post]
func CreatePost(ctx *fiber.Ctx) error {
	var post models.Post

	if err := ctx.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := post.Create(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	//return ctx.Status(fiber.StatusCreated).JSON(post)
	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Create Post successfully",
		Data:    post,
	})
}

// Get Post
// @Summary Get a post
// @Tags /post
// @Accept json
// @Produce json
// @Param id path []byte true "PostID"
// @Success 200 ""
// @Router /post/{postid} [get]
func GetPost(ctx *fiber.Ctx) error {
	var post models.Post
	post_id := ctx.Params("postid")

	result, err := post.GetOne(post_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//return ctx.Status(fiber.StatusCreated).JSON(result)
	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get Post successfully",
		Data:    result,
	})
}

// Get All Posts
// @Summary Get All Posts
// @Tags /post
// @Accept json
// @Produce json
// @Success 200 ""
// @Router /post [get]
func GetPosts(ctx *fiber.Ctx) error {
	var post models.Post

	results, err := post.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//return ctx.Status(fiber.StatusCreated).JSON(results)
	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get Posts successfully",
		Data:    results,
	})
}

// Delete Post
// @Summary Delete a post
// @Tags /post
// @Accept json
// @Produce json
// @Param id path []byte true "PostID"
// @Success 200 ""
// @Router /post/{postid} [delete]
func DeletePost(ctx *fiber.Ctx) error {
	var post models.Post
	post_id := ctx.Params("postid")

	if err := post.DeleteOne(post_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	results, err := post.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//return ctx.SendStatus(200)
	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Delete Post successfully",
		Data:    results,
	})
}

// Update Post
// @Summary Update Post
// @Tags /post
// @Accept json
// @Produce json
// @Param id path []byte true "PostID"
// @Param post body models.Post true "Post"
// @Success 200 ""
// @Router /post/{postid} [put]
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

	result, err := post.GetOne(post_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//return ctx.SendStatus(200)
	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get Post successfully",
		Data:    result,
	})
}
