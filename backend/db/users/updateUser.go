package users

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserPass(c *gin.Context, userName string, password string) bool {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return false
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	// MakeFilterWithNonnilFieldが未完成なので現状ハードコーディング
	// targFilter := db.MakeFilterWithNonnilField(targUser)
	// setFilter := db.MakeFilterWithNonnilField(setInfo)

	_, err := collection.UpdateOne(context.TODO(), bson.M{"userName": userName}, bson.M{"$set": bson.M{"password": password}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return false
	}
	return true
}
