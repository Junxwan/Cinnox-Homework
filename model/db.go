package model

import (
	"Cinnox-Homework/cmd"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type DB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewDB(ctx context.Context, conf cmd.Databases) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Username, conf.Password, conf.Host, conf.Port),
	))

	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &DB{
		client:   client,
		database: client.Database(conf.Databases),
	}, err
}

type User struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserId  string             `bson:"restaurant_id"`
	Name    string
	Created time.Time
}

// 新增用戶資訊
func (db *DB) CreateUser(id, name string, created time.Time) error {
	var user User
	filter := bson.M{"user_id": id}
	coll := db.database.Collection("user")

	if err := coll.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			doc := bson.M{"user_id": id, "name": name, "created": created}

			_, err := coll.InsertOne(context.TODO(), doc)
			return err
		}

		return err
	}

	_, err := coll.UpdateByID(context.TODO(), user.ID, bson.D{
		{"$set", bson.D{{"name", name}}},
	})
	return err
}

// 新增訊息
func (db *DB) CreateMessage(userId, message string, created time.Time) error {
	doc := bson.M{"user_id": userId, "message": message, "created": created}
	_, err := db.database.Collection("message").InsertOne(context.TODO(), doc)
	return err
}

// 關閉連線
func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
