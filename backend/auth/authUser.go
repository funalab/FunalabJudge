package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JsonRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type JsonReturn struct {
	Authorized bool   `json:"authorized"`
	UserName   string `json:"userName"`
	Role       string `json:"role"`
}

type User struct {
	UserId      int64
	Email       string
	Password    string
	CreatedDate time.Time
	Role        string
}

func extractUsername(email string) string {
	parts := strings.Split(email, "@")
	username := parts[0]
	return username
}

func authorizeUser(user User, form JsonRequest) bool {
	return user.Password == form.Password
}

func UserIdInJwt(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	userID := claims[jwt.IdentityKey]
	return userID.(string)
}

func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "funalabJudge",
		Key:            []byte("a2u3zWOTpZKyOkg3NmjVlRnP8x1v4K8KsJv8NDFlTSY="), // TODO .envファイルに移動, 運用時には再作成: % openssl rand -base64 32
		Timeout:        time.Hour * 24 * 7,                                     // equals to CookieMaxAge
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

			var result User
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			client, _ := mongo.Connect(
				ctx,
				options.Client().ApplyURI("mongodb://localhost:27017/"),
			)
			err := client.Database("dev").Collection("users").FindOne(context.TODO(), bson.D{{"email", jsonRequest.UserId}}).Decode(&result)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return "", jwt.ErrMissingLoginValues
			}
			if result.Password != jsonRequest.Password {
				return "", jwt.ErrFailedAuthentication
			}

			return result.Email, nil
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
