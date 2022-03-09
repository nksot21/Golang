package handler

import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"
	"errors"

	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserToken struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, Name: user.Name, Email: user.Email}
}

func CreateResponseUserToken(user models.User, token string, exp int64) UserToken {
	return UserToken{ID: user.ID, Name: user.Name, Email: user.Email, Token: token, Exp: exp}
}

// CHECK EMAIL EXISTED
func CheckUserEmail(email string) error {
	userCol := firebase.FirebaseApp.Db.Collection("users")
	query := userCol.Where("email", "==", email)

	_, err := query.Documents(firebase.Ctx).Next()
	if err == nil {
		return errors.New("User existed")
	}
	return nil
}

func findUserEmail(email string) (models.User, error) {
	var user models.User
	userCol := firebase.FirebaseApp.Db.Collection("users")
	query := userCol.Where("email", "==", email)

	userDB, err := query.Documents(firebase.Ctx).Next()
	if err != nil {
		return user, errors.New("User not exist")
	}
	if err = userDB.DataTo(&user); err != nil {
		return user, err
	}
	user.ID = userDB.Ref.ID
	return user, nil
}

// CREATE JWT TOKEN
func createJWTToken(user *models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("privatekey"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

// Create New User
// @Summary      Create a new user
// @Description  Create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        todo  body      types.User  true  "New Todo"
// @Success      200   {object}  types.User
// @Failure      400   {object}  HTTPError
// @Router       /user [post]
func SignUp(c *fiber.Ctx) error {
	var user models.User
	userCol := firebase.FirebaseApp.Db.Collection("users")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if user.Email == "" || user.Name == "" || user.Password == "" {
		return c.Status(400).JSON("error email/ name/ password")
	}

	err := CheckUserEmail(user.Email)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	user.Password = string(hash)

	newUser := userCol.NewDoc()
	user.ID = newUser.ID
	_, err = newUser.Create(firebase.Ctx, user)
	if err != nil {
		return c.Status(400).JSON(err)
	}
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

func GetUserByID(userID string) (*firestore.DocumentRef, error) {
	userCol := firebase.FirebaseApp.Db.Collection("users")
	user := userCol.Doc(userID)
	_, err := user.Get(firebase.Ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
