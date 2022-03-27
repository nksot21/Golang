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
	BaseModel `bson:",inline"`

	ID             primitive.ObjectID `json:"id,omitempty" bson:"id"`
	Title          string             `json:"title,omitempty" bson:"title"`
	Emotion        int                `json:"emotion,omitempty" bson:"emotion"`
	Detail         string             `json:"detail,omitempty" bson:"detail"`
	Picture        string             `json:"picture,omitempty" bson:"picture"`
	FireBaseUserId string             `json:"firebase_user_id,omitempty" bson:"firebase_user_id"`
	Expert         User               `json:"expert" bson:"expert"`
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

	p.FireBaseUserId = newPost.FireBaseUserId

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

	//fmt.Println(post.ID)
	if err != nil {
		return post, err
	}

	var user User
	user.GetOne(post.FireBaseUserId, "")
	post.Expert = user

	return post, nil
}

func (p *Post) GetAll() ([]Post, error) {
	c, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"deleted": false,
	}

	instance := database.GetMongoInstance()
	results, err := instance.Db.Collection("Posts").Find(c, filter)

	//fmt.Println(results)

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

		var user User
		user.GetOne(post.FireBaseUserId, "")
		post.Expert = user

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
	objId, _ := primitive.ObjectIDFromHex(post_id)
	/*post, err := p.GetOne(post_id)

	if err != nil {
		return err
	}

	newBaseModule := BaseModel{
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Deleted:   true,
		DeletedAt: time.Now().Unix(),
	}*/

	update := bson.M{
		"$set": bson.M{
			"deleted":    true,
			"deleted_at": time.Now().Unix(),
		},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"id": objId}, update)
	fmt.Println(result)
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

	post, err := p.GetOne(post_id)

	if err != nil {
		return err
	}

	newBaseModule := BaseModel{
		CreatedAt: post.CreatedAt,
		UpdatedAt: time.Now().Unix(),
		Deleted:   post.Deleted,
		DeletedAt: post.DeletedAt,
	}

	updatePost := bson.M{
		"title":            p.Title,
		"emotion":          p.Emotion,
		"detail":           p.Detail,
		"picture":          p.Picture,
		"firebase_user_id": p.FireBaseUserId,
		"basemodel":        newBaseModule,
	}

	instance := database.GetMongoInstance()
	result, err := instance.Db.Collection("Posts").UpdateOne(context.Background(), bson.M{"id": objId}, bson.M{"$set": updatePost})
	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}
