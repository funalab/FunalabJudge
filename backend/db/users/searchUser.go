package users

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchOneUserWithUserName(client *mongo.Client, userName string) (User, error) {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	filter := bson.M{"userName": userName}

	var u User
	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	return u, err
}

func SearchUsersWithJoinnedDate(client *mongo.Client, st time.Time, ed time.Time) ([]string, error) {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)
	filter := bson.M{"joinedDate": bson.M{
		"$gte": st,
		"$lte": ed,
	},
	}
	var us []User
	cur, err := collection.Find(context.TODO(), filter)
	if err = cur.All(context.TODO(), &us); err != nil {
		return nil, err
	}
	var userNames []string
	for _, u := range us {
		userNames = append(userNames, u.UserName)
	}

	return userNames, nil

}
