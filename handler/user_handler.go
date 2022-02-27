package handler

import (
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("firebase_user_id")
	var user models.User
	err := user.GetOne(id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := user.Create(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	firebaseUserId := ctx.Params("firebase_user_id")
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := user.Update(firebaseUserId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func AddFeelUser(ctx *fiber.Ctx) error {
	firebaseUserId := ctx.Params("firebase_user_id")
	var user models.User
	var feel models.UserFeel

	if err := ctx.BodyParser(&feel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.GetOne(firebaseUserId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.AddUserFeel(feel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
