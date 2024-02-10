package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ConnectionErr struct {
	msg string
}

func (c *ConnectionErr) Error() string {
	return fmt.Sprintf("ConnectionErr: %v\n", c.msg)
}

func NewConnectionErr(errMsg string) error {
	return &ConnectionErr{
		msg: errMsg,
	}
}

func Mongo_connectable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	if err != nil {
		log.Fatalf("connection error :%v", err)
		return false
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return false
	}
	cancel()
	if err := mongoClient.Disconnect(ctx); err != nil {
		log.Fatalf("mongodb disconnect error : %v", err)
		return false
	}
	return true
}
