package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() *Mongo {

	connStr := "mongodb://root:tpwd@localhost:27017"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return &Mongo{
		client: client,
	}
}

type Mongo struct {
	client *mongo.Client
}

func (m *Mongo) Disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("Disconnected")
}

func (m *Mongo) Ping() {

	var result bson.M
	bsonQuery := bson.D{{
		"ping",
		1,
	}}

	db := m.client.Database("admin")
	iresult := db.RunCommand(context.TODO(), bsonQuery)

	if err := iresult.Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}
