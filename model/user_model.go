package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mental-health-api/pkg/const/collections"
	"mental-health-api/pkg/database"
	"time"
)

type User struct {
	BaseModel `bson:",inline"`

	ID             primitive.ObjectID `json:"-" bson:"_id, omitempty"`
	FireBaseUserId string             `json:"firebase_user_id" bson:"fire_base_user_id, omitempty"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	Bio            string             `json:"bio" bson:"bio"`
	IsExpert       bool               `json:"is_expert" bson:"is_expert"`
}

func (u *User) GetOne(firebaseUserId string) error {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := map[string]string{
		"fire_base_user_id": firebaseUserId,
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	return err
}

func (u *User) Create() error {
	u.BaseModel = BaseModel{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	u.ID = primitive.NewObjectID()

	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	user := collection.FindOne(context.TODO(), bson.M{
		"fire_base_user_id": u.FireBaseUserId,
	})

	if user.Err() == nil {
		return errors.New("User already exists")
	}

	_, err := collection.InsertOne(context.Background(), u)
	return err
}

func (u *User) Update(firebaseUserId string) error {
	u.BaseModel.UpdatedAt = time.Now().Unix()
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{"fire_base_user_id": firebaseUserId}

	update := bson.M{
		"$set": bson.M{
			"name":       u.Name,
			"email":      u.Email,
			"bio":        u.Bio,
			"is_expert":  u.IsExpert,
			"updated_at": u.BaseModel.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (u *User) Delete(id string) error {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
			"deleted":    true,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (u *User) AddUserFeel(userFeel UserFeel) error {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_COLLECTION)
	filter := bson.M{
		"_id": u.ID.Hex(),
	}
	update := bson.M{
		"$push": map[string]interface{}{
			"user_feels": userFeel,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err

}
