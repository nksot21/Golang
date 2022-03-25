package models

import (
	"context"
	"errors"
	"fmt"
	"mental-health-api/pkg/const/collections"
	"mental-health-api/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseModel `bson:",inline"`

	ID             primitive.ObjectID `json:"-" bson:"_id, omitempty"`
	FireBaseUserId string             `json:"firebase_user_id" bson:"firebase_user_id, omitempty"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	Bio            string             `json:"bio" bson:"bio"`
	Picture        string             `json:"picture" bson:"picture"`
	IsExpert       bool               `json:"is_expert" bson:"is_expert"`
}

func (u *User) GetOne(firebaseUserId string, email string) error {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"firebase_user_id": firebaseUserId,
		"email":            email,
		"deleted":          false,
	}

	if email != "" {
		delete(filter, "firebase_user_id")
	} else {
		delete(filter, "email")
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	return err
}

func (u *User) Create(checkExist bool) error {
	u.BaseModel = BaseModel{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	u.ID = primitive.NewObjectID()

	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)

	if checkExist == true {
		user := collection.FindOne(context.TODO(), bson.M{
			"firebase_user_id": u.FireBaseUserId,
			"deleted":          false,
		})
		if user.Err() == nil {
			return errors.New("User already exists")
		}
	}

	if u.Picture == "" {
		u.Picture = "https://images.pexels.com/photos/9456631/pexels-photo-9456631.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940"
	}

	_, err := collection.InsertOne(context.Background(), u)
	return err
}

func (u *User) Update(firebaseUserId string) error {
	u.BaseModel.UpdatedAt = time.Now().Unix()
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"firebase_user_id": firebaseUserId,
		"deleted":          false,
	}

	update := bson.M{
		"$set": bson.M{
			"name":       u.Name,
			"bio":        u.Bio,
			"picture":    u.Picture,
			"is_expert":  u.IsExpert,
			"updated_at": u.BaseModel.UpdatedAt,
		},
	}

	result := collection.FindOneAndUpdate(context.Background(), filter, update)
	return result.Err()
}

func (u *User) Delete(firebaseUserId string) error {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"firebase_user_id": firebaseUserId,
		"deleted":          false,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
			"deleted":    true,
		},
	}
	result := collection.FindOneAndUpdate(context.Background(), filter, update)

	return result.Err()
}

func (u *User) GetAll() ([]User, error) {
	c, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"deleted": false,
	}

	results, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	defer results.Close(c)

	var users []User
	for results.Next(c) {
		var user User
		if err := results.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	fmt.Println(users)

	return users, nil
}
