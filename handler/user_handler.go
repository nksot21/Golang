package handler

import (
	"fmt"
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.Create(false); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User login successfully",
		Data:    user,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}

	email := ctx.Query("email")
	fmt.Println(firebaseid, email)

	var user models.User
	err := user.GetOne(firebaseid, email)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get User info successfully",
		Data:    user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := user.Create(true); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := user.Update(firebaseid); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}

	var user models.User

	if err := user.Delete(firebaseid); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
