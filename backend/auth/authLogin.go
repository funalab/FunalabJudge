package auth

import (
	"go-test/db/users"
	"go-test/util"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func LoginAuthenticator(c *gin.Context) (interface{}, error) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
		return "", jwt.ErrFailedAuthentication
	}
	client := client_.(*mongo.Client)

	var jsonRequest LoginRequest
	if err := c.ShouldBind(&jsonRequest); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	u, err := users.SearchUserWithUserName(client, jsonRequest.UserName)
	if err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	if !CheckPasswordHash(jsonRequest.Password, u.Password) {
		return "", jwt.ErrFailedAuthentication
	}

	return &u, nil
}

func JwtMapper(data interface{}) jwt.MapClaims {
	// 引数"data"はLoginAuthenticatorの一つ目のreturn
	if v, ok := data.(*users.User); ok {
		return jwt.MapClaims{
			JwtIdentityKey: v.UserName,
			JwtUserRoleKey: v.Role,
		}
	}
	return jwt.MapClaims{}
}
