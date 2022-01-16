package router

import (
	"DBdemo/src/model"

	"github.com/gofiber/fiber/v2"
)

type Client struct {
	ID     int
	Name   string
	Age    string
	Gender string
	Email  string
}

func CreateResponseUser(client model.Client) Client {
	return Client{ID: client.ID, Name: client.Name, Age: client.Age, Gender: client.Gender, Email: client.Email}
}

func CreateClient(c *fiber.Ctx) error {
	//ClientRepo := implement.NewClientRepo(db.SQL)
	var client model.Client
	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//ClientRepo.Insert(client)
	responseClient := CreateResponseUser(client)

	return c.Status(200).JSON(responseClient)
}
