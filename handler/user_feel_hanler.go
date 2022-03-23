package handler

import (
	"fmt"
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUserFeel(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}

	var user models.User

	if err := user.GetOne(firebaseid, ""); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var userFeel models.UserFeel
	if err := ctx.BodyParser(&userFeel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := userFeel.Create(firebaseid); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User Feel created successfully",
		Data:    userFeel,
	})
}

func GetUserFeel(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")

	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}

	var userFeel models.UserFeel
	data, err := userFeel.GetFeels(firebaseid)

	fmt.Println(data)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get User info successfully",
		Data:    &data,
	})
}
