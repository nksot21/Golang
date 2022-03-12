package handler

import (
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

// Create Post
// @Summary Create Post
// @Description EVENT EMOTION = 0
// @Description POST  HAPPY = 1
// @Description POST  SAD = 2
// @Description POST  SCARED = 3
// @Description POST  ANGRY = 4
// @Description POST  WORRY = 5
// @Description POST  NORMAL = 6
// @Description POST  DEPRESSION = 7
// @Tags /post
// @Accept json
// @Produce json
// @Param post body models.Post true "Post"
// @Success 200 ""
// @Router /post/create [post]
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

	return ctx.Status(fiber.StatusCreated).JSON(result)
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

	return ctx.Status(fiber.StatusCreated).JSON(results)
}

// Delete Post
// @Summary Delete a post
// @Tags /post
// @Accept json
// @Produce json
// @Param id path []byte true "PostID"
// @Success 200 ""
// @Router /post/delete/{postid} [put]
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

	if err := post.DeleteOne(post_id); err != nil {
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

// Update Post
// @Summary Update Post
// @Tags /post
// @Accept json
// @Produce json
// @Param id path []byte true "PostID"
// @Param post body models.Post true "Post"
// @Success 200 ""
// @Router /post/update/{postid} [put]
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
