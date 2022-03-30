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

	ID             primitive.ObjectID `json:"id" bson:"_id"`
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

	newFeel := UserFeel{
		BaseModel: BaseModel{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		ID:             primitive.NewObjectID(),
		FeelId:         uf.FeelId,
		Reason:         uf.Reason,
		FireBaseUserId: firebaseUserId,
	}

	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)
	_, err := collection.InsertOne(context.TODO(), newFeel)
	return err
}

func (uf *UserFeel) GetFeels(firebaseUserId string) ([]UserFeel, error) {
	uf.BaseModel.UpdatedAt = time.Now().Unix()
	collection := database.GetMongoInstance().Db.Collection(collections.USER_FEEL_COLLECTION)

	filter := bson.M{
		"firebase_user_id": firebaseUserId,
	}

	filterOption := options.Find()
	filterOption.SetSort(bson.M{
		"created_at": -1,
	})

	var results []UserFeel
	cur, err := collection.Find(context.TODO(), filter, filterOption)

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem UserFeel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)

	}

	return results, err
}
