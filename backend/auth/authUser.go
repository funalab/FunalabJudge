package auth

import (
	"go-test/db/submission"
	"go-test/db/users"
	"go-test/util"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserFromJwt(c *gin.Context) interface{} {
	// jwtからuser情報を抽出し、UserAuthorizatorに渡す
	claims := jwt.ExtractClaims(c)
	joinedDate, _ := time.Parse(time.RFC3339, claims[JwtJoinedDateKey].(string))
	return &users.User{
		UserName:   claims[JwtIdentityKey].(string),
		JoinedDate: joinedDate.Local(),
	}
}

func UserAuthorizator(data interface{}, c *gin.Context) bool {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
		return false
	}
	client := client_.(*mongo.Client)
	// 引数"data"はGetUserNameFromJwtのreturn
	if jwtUser, ok := data.(*users.User); ok {
		if jwtUser.JoinedDate.Year() < time.Now().Local().Year() {
			return true
		} else {
			urlUserName := c.Param("userName")
			if urlUserName == "" { // userNameを含まないエンドポイントの場合
				if c.Param("submissionId") != "" {
					sId, err := primitive.ObjectIDFromHex(c.Param("submissionId"))
					if err != nil {
						log.Printf("Failed to parse submissionId : %s\n", err.Error())
					}
					s, err := submission.SearchOneSubmissionWithId(client, sId)
					if err != nil {
						log.Printf("Failed to search submission from id : %s\n", err.Error())
					}
					u, err := users.SearchOneUserWithUserName(client, s.UserName)
					if err != nil {
						log.Printf("Failed to search user from userName : %s\n", err.Error())
					}
					urlUserName = u.UserName
				} else {
					// userNameもsubmissionIdもない = 全ユーザがアクセス可能なエンドポイント
					return true
				}
			}
			if jwtUser.UserName == urlUserName {
				return true
			}
		}
	}

	return false
}
