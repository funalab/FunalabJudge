package auth

import (
	"context"
	"go-test/types"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JsonRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type JsonReturn struct {
	Authorized bool   `json:"authorized"`
	UserName   string `json:"userName"`
	Role       string `json:"role"`
}

func UserIdInJwt(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	userID := claims[jwt.IdentityKey]
	return userID.(string)
}

func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "funalabJudge",
		Key:            []byte(os.Getenv("SECRET_KEY")), // 運用時には再作成する: % openssl rand -base64 32
		Timeout:        time.Hour * 24 * 7,              // equals to CookieMaxAge
		MaxRefresh:     time.Hour * 24 * 7,
		SendCookie:     true,
		SecureCookie:   false, //non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieDomain:   "localhost:3000",
		CookieName:     "token", // default jwt
		TokenLookup:    "cookie:token",
		CookieSameSite: http.SameSiteStrictMode, //SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode
		IdentityKey:    "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var jsonRequest JsonRequest

			if err := c.ShouldBind(&jsonRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			dbName := os.Getenv("DB_NAME")
			usrCol := os.Getenv("USERS_COLLECTION")
			client, _ := c.Get("mongoClient")
			dbClient := client.(*mongo.Client)
			filter := bson.M{"userName": jsonRequest.UserName}

			var result types.User
			err := dbClient.Database(dbName).Collection(usrCol).FindOne(context.TODO(), filter).Decode(&result)
			if err != nil {
				println(err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			if result.Password != jsonRequest.Password {
				return "", jwt.ErrFailedAuthentication
			}

			return result.UserName, nil
		},
	})

	if err != nil {
		return nil, err
	}

	errInit := jwtMiddleware.MiddlewareInit()
	if errInit != nil {
		return nil, err
	}

	return jwtMiddleware, nil
}
