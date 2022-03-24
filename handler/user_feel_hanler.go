package handler

import (
	"fmt"
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

// Create User's Feeling
// @Summary Create User's Feel
// @Description EVENT_emotion = 0
// @Description POST__happy = 1
// @Description POST__sad = 2
// @Description POST__scared = 3
// @Description POST__angry = 4
// @Description POST__worry = 5
// @Description POST__normal = 6
// @Description POST__depression = 7
// @Tags /user-feel
// @Accept json
// @Produce json
// @Param UserFeel body models.UserFeel true "User Feel"
// @Success 200 ""
// @Router /user-feel [post]
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

// Get User's Feeling
// @Summary Get User's Feel
// @Description EVENT_emotion = 0
// @Description POST__happy = 1
// @Description POST__sad = 2
// @Description POST__scared = 3
// @Description POST__angry = 4
// @Description POST__worry = 5
// @Description POST__normal = 6
// @Description POST__depression = 7
// @Tags /user-feel
// @Accept json
// @Produce json
// @Param UserID header string true "UserID"
// @Success 200 ""
// @Router /user-feel [get]
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
