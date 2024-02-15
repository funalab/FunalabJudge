package auth

import (
	"context"
	"go-test/types"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JsonRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func LoginAuthenticator(c *gin.Context) (interface{}, error) {
	var jsonRequest JsonRequest

	if err := c.ShouldBind(&jsonRequest); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	client, _ := c.Get("mongoClient")
	dbClient := client.(*mongo.Client)

	var result types.User
	filter := bson.M{"userName": jsonRequest.UserName}
	err := dbClient.Database(dbName).Collection(usrCol).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	if result.Password != jsonRequest.Password {
		return "", jwt.ErrFailedAuthentication
	}

	return &result, nil
}

func JwtMapper(data interface{}) jwt.MapClaims {
	// 引数"data"はLoginAuthenticatorの一つ目のreturn
	if v, ok := data.(*types.User); ok {
		return jwt.MapClaims{
			JwtIdentityKey: v.UserName,
			JwtUserRoleKey: v.Role,
		}
	}
	return jwt.MapClaims{}
}