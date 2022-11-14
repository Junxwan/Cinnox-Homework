package model

import (
	"Cinnox-Homework/cmd"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
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
		client: client,
	}, err
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
