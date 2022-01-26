package handler

import (
	"connectdb/src/driver"
	"connectdb/src/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserToken struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

// CREATE RESPONSE FUNCTION
func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, Name: user.Name, Email: user.Email}
}

func CreateResponseUserToken(user models.User, token string, exp int64) UserToken {
	return UserToken{ID: user.ID, Name: user.Name, Email: user.Email, Token: token, Exp: exp}
}

// CHECK USER BY ID
func findUserID(id int, user *models.User) error {
	driver.PostgresDB.SQL.Find(&user, "id=?", id)

	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

// GET USER BY EMAIL
func findUserEmail(email string) (models.User, error) {
	var user models.User
	driver.PostgresDB.SQL.Find(&user, "email=?", email)
	if user.ID == 0 {
		return user, errors.New("User does not exist")
	}
	return user, nil
}

// CHECK EMAIL EXISTED
func checkUserEmail(email string) error {
	var user models.User
	driver.PostgresDB.SQL.Find(&user, "email=?", email)
	if user.ID != 0 {
		return errors.New("User existed! Try another email")
	}
	return nil
}

// CREATE JWT TOKEN
func createJWTToken(user *models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

// CREATE USER: POST: user/new
func SignUp(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if user.Email == "" || user.Name == "" || user.Password == "" {
		return c.Status(400).JSON("error email/ name/ password")
	}

	err := checkUserEmail(user.Email)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	user.Password = string(hash)

	driver.PostgresDB.SQL.Create(&user)
	token, exp, err := createJWTToken(&user)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	responseUser := CreateResponseUserToken(user, token, exp)

	return c.Status(200).JSON(responseUser)
}

func SignIn(c *fiber.Ctx) error {
	var reqUser models.User
	if err := c.BodyParser(&reqUser); err != nil {
		return c.Status(400).JSON("error")
	}

	userdb, err := findUserEmail(reqUser.Email)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userdb.Password), ([]byte(reqUser.Password))); err != nil {
		return c.Status(400).JSON("Wrong Password")
	}

	token, exp, err := createJWTToken(&reqUser)
	if err != nil {
		return c.Status(400).JSON("error")
	}

	responseUser := CreateResponseUserToken(userdb, token, exp)
	return c.Status(200).JSON(responseUser)
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
	if err := findUserID(id, &user); err != nil {
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

	if err := findUserID(id, &user); err != nil {
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
	if err := findUserID(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := driver.PostgresDB.SQL.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Deleted User")
}
