package routes

import (
	"connectdb/src/driver"
	"connectdb/src/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	OrderDate time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, OrderDate: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON("error bodyparser")
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON("error find user")
	}
	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON("error find product")
	}

	driver.PostgresDB.SQL.Create(&order)
	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), createProductResponse(product))
	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	driver.PostgresDB.SQL.Find(&orders)

	responseOrders := []Order{}
	for _, order := range orders {
		var user User
		var product Product
		driver.PostgresDB.SQL.Find(&user, "id=?", order.UserRefer)
		driver.PostgresDB.SQL.Find(&product, "id=?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, user, product)
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	driver.PostgresDB.SQL.Find(&order, "id=?", id)

	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	var order models.Order
	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product
	driver.PostgresDB.SQL.First(&user, order.UserRefer)
	driver.PostgresDB.SQL.First(&product, order.ProductRefer)

	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), createProductResponse(product))
	return c.Status(200).JSON(responseOrder)
}
