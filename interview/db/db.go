package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Con *mongo.Database
)

const (
	MongoURI = "mongodb://localhost:27017/test"
)

func ConnectDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opt := options.Client().ApplyURI(MongoURI)
	c, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	Con = c.Database("test")
}

type Image struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Header   []string           `bson:"header,omitempty"`
	Size     int64              `bson:"size,omitempty"`
	Filename string             `bson:"filename,omitempty"`
}

func (i *Image) Insert() error {
	c := Con.Collection("image")
	r, err := c.InsertOne(context.TODO(), i)
	if err != nil {
		return err
	}
	fmt.Println(r.InsertedID)
	return nil
}

