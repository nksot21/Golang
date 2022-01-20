package routes

import (
	"connectdb/src/driver"
	"connectdb/src/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func createProductResponse(product models.Product) Product {
	return Product{ID: product.ID, Name: product.Name}
}

// CREATE PRODUCT
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	driver.PostgresDB.SQL.Create(&product)
	responseProduct := createProductResponse(product)
	return c.Status(200).JSON(responseProduct)
}

// GET ALL
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	if err := driver.PostgresDB.SQL.Find(&products).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := createProductResponse(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

// CHECK ID
func findProduct(id int, product *models.Product) error {
	driver.PostgresDB.SQL.Find(&product, "id=?", id)

	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

// GET BY ID
func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := createProductResponse(product)
	return c.Status(200).JSON(responseProduct)
}

// UPDATE PRODUCT
func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateProduct struct {
		Name string `json:"name"`
	}

	var updateData updateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.Name = updateData.Name
	driver.PostgresDB.SQL.Save(&product)

	responseProduct := createProductResponse(product)
	return c.Status(200).JSON(responseProduct)
}

// DELETE PRODUCT
func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("wrong ID")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := driver.PostgresDB.SQL.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Deleted Product")
}
