package handler

import (
	"connectdb/src/driver"
	"connectdb/src/models"
	"errors"

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
	driver.PostgresDB.SQL.Find(&user, "id=?", id)

	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

// GET ALL USERS: GET: /user
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	driver.PostgresDB.SQL.Find(&users)

	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

// GET USER BY ID: GET: /user/:id
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

// UPDATE USER: PUT: /user/:id
func UpdateUser(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//update struct
	type updateUser struct {
		Name string `json:"name"`
	}
	var updateData updateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.Name = updateData.Name
	driver.PostgresDB.SQL.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

// DELETE USER: DELETE: /user/:id
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong id")
	}

	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := driver.PostgresDB.SQL.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Deleted User")
}
