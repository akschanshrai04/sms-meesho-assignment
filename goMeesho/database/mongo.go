package database

import (
	"goMeesho/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.Connect(options.Client().ApplyURI(config.GetMongoURI()))

	if err != nil {
		panic(err)
	}

	DB = client.Database(config.GetMongoDBName())
}
