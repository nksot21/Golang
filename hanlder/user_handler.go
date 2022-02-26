package hanlder

import (
	"github.com/gofiber/fiber/v2"
	models "mental-health-api/model"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if _, err := user.Create(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func CrateExpert() {

}
