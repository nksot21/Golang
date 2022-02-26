package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mental-health-api/pkg/database"
)

type User struct {
	BaseModel

	FireBaseUserId string `json:"firebase_user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Bio            string `json:"bio"`
}

func (u *User) Create() (*mongo.InsertOneResult, error) {
	fmt.Println(u)
	instance := database.GetMongoInstance()
	result, err := instance.Db.Collection("users").InsertOne(context.Background(), u)
	fmt.Println(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
