package models

import (
	"context"
	"mental-health-api/pkg/const/collections"
	"mental-health-api/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserFeel struct {
	BaseModel `bson:",inline"`

	ID             primitive.ObjectID `json:"-" bson:"_id"`
	FireBaseUserId string             `json:"firebase_user_id" bson:"firebase_user_id"`

	FeelId int    `json:"feel_id" bson:"feel_id"`
	Reason string `json:"reason" bson:"reason"`
}

func (uf *UserFeel) Create(firebaseUserId string) error {
	uf.BaseModel = BaseModel{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	uf.ID = primitive.NewObjectID()
	uf.FireBaseUserId = firebaseUserId

	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)
	_, err := collection.InsertOne(context.TODO(), uf)
	return err
}

func (uf *UserFeel) GetFeels(firebaseUserId string) error {
	uf.BaseModel.UpdatedAt = time.Now().Unix()
	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)

	filter := bson.M{
		"firebase_user_id": firebaseUserId,
	}

	filterOption := options.Find()
	filterOption.SetSort(bson.M{
		"created_at": -1,
	})

	_, err := collection.Find(context.TODO(), filter, filterOption)

	return err
}
