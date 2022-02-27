package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mental-health-api/pkg/const/collections"
	"mental-health-api/pkg/database"
	"time"
)

type UserFeel struct {
	BaseModel `bson:",inline"`

	ID             primitive.ObjectID `json:"-" bson:"_id"`
	UserId         int                `json:"user_id" bson:"user_id"`
	UserFirebaseId string             `json:"user_firebase_id" bson:"user_firebase_id"`
	FeelId         int                `json:"feel_id" bson:"feel_id"`
}

func (uf *UserFeel) Create() error {
	uf.BaseModel = BaseModel{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	uf.ID = primitive.NewObjectID()

	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)
	_, err := collection.InsertOne(context.TODO(), uf)
	return err
}

func (uf *UserFeel) GetLastUserFeel(firebaseUserId string) {
	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)
	filter := bson.M{"user_firebase_id": firebaseUserId}
	err := collection.FindOne(context.TODO(), filter).Decode(&uf)
	if err != nil {
		panic(err)
	}
}
