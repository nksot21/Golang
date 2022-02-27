package models

import (
	"context"
	"fmt"
	"mental-health-api/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	BaseModel

	ID             primitive.ObjectID `json: "id,omitempty"`
	Title          string             `json: "title,omitempty"`
	Emotion        int                `json: "emotion,omitempty"`
	Detail         string             `json: "detail,omitempty"`
	Picture        string             `json: "picture,omitempty"`
	FireBaseUserId string             `json: "firebase_user_id,omitempty"`
}

type JsonData struct {
	Data []Post `json:"data"`
}

func (p *Post) Create() error {
	p.BaseModel = BaseModel{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	newPost := Post{
		BaseModel: BaseModel{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		ID:             primitive.NewObjectID(),
		Title:          p.Title,
		Emotion:        p.Emotion,
		Detail:         p.Detail,
		Picture:        p.Picture,
		FireBaseUserId: p.FireBaseUserId,
	}

	instance := database.GetMongoInstance()
	result, err := instance.Db.Collection("Posts").InsertOne(context.Background(), newPost)
	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}

func (p *Post) GetOne(post_id string) (Post, error) {
	var post Post
	objId, _ := primitive.ObjectIDFromHex(post_id)

	instance := database.GetMongoInstance()
	err := instance.Db.Collection("Posts").FindOne(context.Background(), bson.M{"id": objId}).Decode(&post)

	fmt.Println(post.ID)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (p *Post) GetAll() ([]Post, error) {
	c, _ := context.WithTimeout(context.Background(), 10*time.Second)

	instance := database.GetMongoInstance()
	results, err := instance.Db.Collection("Posts").Find(c, bson.M{})

	fmt.Println(results)

	if err != nil {
		return nil, err
	}
	defer results.Close(c)

	var posts []Post
	for results.Next(c) {
		var post Post
		if err := results.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	fmt.Println(posts)

	return posts, nil
}

func (p *Post) DeleteOne(post_id string) error {
	/*objId, _ := primitive.ObjectIDFromHex(post_id)

	instance := database.GetMongoInstance()
	results, err := instance.Db.Collection("Posts").DeleteOne(context.Background(), bson.M{"id": objId})

	fmt.Println(results)

	if err != nil {
		return err
	}

	return nil*/
	collection := database.GetMongoInstance().Db.Collection("Posts")
	filter := map[string]string{
		"_id": post_id,
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

func (p *Post) DeleteAll() error {
	instance := database.GetMongoInstance()
	_, err := instance.Db.Collection("Posts").DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (p *Post) UpdateOne(post_id string) error {
	objId, _ := primitive.ObjectIDFromHex(post_id)

	updatePost := bson.M{
		"title":            p.Title,
		"emotion":          p.Emotion,
		"detail":           p.Detail,
		"picture":          p.Picture,
		"firebase_user_id": p.FireBaseUserId,
	}

	instance := database.GetMongoInstance()
	result, err := instance.Db.Collection("Posts").UpdateOne(context.Background(), bson.M{"id": objId}, bson.M{"$set": updatePost})
	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}
