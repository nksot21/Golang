package handler

import (
	"fmt"
	models "mental-health-api/model"

	"github.com/gofiber/fiber/v2"
)

// Login
// @Summary Login
// @Tags /user
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user/login [post]
func Login(ctx *fiber.Ctx) error {
	firebaseid := ctx.Get("x-firebase-uid")
	if firebaseid == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "header.x-firebase-uid is empty")
	}
	/*var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.Create(false); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}*/

	email := ctx.Query("email")
	fmt.Println(firebaseid, email)

	var user models.User
	err := user.GetOne(firebaseid, email)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "User login successfully",
		Data:    user,
	})
}

// Get User
// @Summary Get User
// @Tags /user
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user/get-info [get]
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

// Create User
// @Summary Create User
// @Tags /user
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.Response
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var checkUser models.User

	if err := checkUser.GetOne(user.FireBaseUserId, ""); err == nil {
		return ctx.JSON(models.Response{
			Status:  400,
			Error:   true,
			Message: "User already exists in the database",
		})
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

// Update User
// @Summary Update User
// @Tags /user
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /user [put]
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

// Delete User
// @Summary Delete User
// @Tags /user
// @Accept json
// @Produce json
// @Param userID header string true "UserID"
// @Success 200 ""
// @Router /user [delete]
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

func GetUsers(ctx *fiber.Ctx) error {
	var user models.User
	results, err := user.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(models.Response{
		Status:  fiber.StatusCreated,
		Message: "Get User info successfully",
		Data:    results,
	})
}
