package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mental-health-api/pkg/utils"
	"os"
	"sync"
)

type IMongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var instance *IMongoInstance
var one sync.Once

func GetMongoInstance() *IMongoInstance {
	one.Do(func() {
		mongoUrl, err := utils.ConnectionURLBuilder("mongodb")
		if err != nil {
			panic(err)
		}

		clientOptions := options.Client().ApplyURI(mongoUrl)
		client, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			panic(err)
		}

		instance = &IMongoInstance{
			Client: client,
			Db:     client.Database(os.Getenv("MONGO_DBNAME")),
		}
	})

	return instance
}
