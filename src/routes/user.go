package routes

import (
	"connectdb/src/driver"
	"connectdb/src/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CREATE RESPONSE FUNCTION
func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, Name: user.Name}
}

// CREATE USER: POST: /newuser
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	driver.PostgresDB.SQL.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// CHECK USER BY ID
func findUser(id int, user *models.User) error {
	return nil
}

// GET ALL USERS: GET: /user
func GetUsers(c *fiber.Ctx) error {
	return nil
}

// GET USER BY ID: GET: /user/:id
func GetUser(c *fiber.Ctx) error {
	return nil
}

// UPDATE USER: PUT: /user/:id
func UpdateUser(c *fiber.Ctx) error {
	return nil
}

// DELETE USER: DELETE: /user/:id
func DeleteUser(c *fiber.Ctx) error {
	return nil
}
