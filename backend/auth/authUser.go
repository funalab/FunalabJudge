package auth

import (
	"fmt"
	"go-test/db/submission"
	"go-test/db/users"
	"go-test/util"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserNameFromJwt(c *gin.Context) interface{} {
	// jwtからuser情報を抽出し、UserAuthorizatorに渡す
	claims := jwt.ExtractClaims(c)
	return &users.User{
		UserName: claims[JwtIdentityKey].(string),
		Role:     claims[JwtUserRoleKey].(string),
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
		if jwtUser.Role == "admin" || jwtUser.Role == "manager" {
			return true
		} else if jwtUser.Role == "user" {
			urlUserName := c.Param("userName")
			if urlUserName == "" { // userNameを含まないエンドポイントの場合
				if c.Param("submissionId") != "" {
					urlSubmissionId, err := strconv.Atoi(c.Param("submissionId"))
					if err != nil {
						fmt.Println(err)
						return false
					}
					s, err := submission.SearchSubmissionWithId(client, int32(urlSubmissionId))
					u, err := users.SearchUserWithUserName(client, s.UserName)
					urlUserName = u.UserName
				} else {
					// userNameもsubmissionIdもない = 全ユーザがアクセス可能なエンドポイント
					return true
				}
			}
			if jwtUser.UserName == urlUserName {
				return true
			}
		} else {
			// unexpected Role
			return false
		}
	}

	return false
}
