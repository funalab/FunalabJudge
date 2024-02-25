package auth

import (
	"fmt"
	"go-test/db/submission"
	"go-test/db/users"
	"go-test/myTypes"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserNameFromJwt(c *gin.Context) interface{} {
	// jwtからuser情報を抽出し、UserAuthorizatorに渡す
	claims := jwt.ExtractClaims(c)
	return &myTypes.User{
		UserName: claims[JwtIdentityKey].(string),
		Role:     claims[JwtUserRoleKey].(string),
	}
}

func GetUserNameFromsubmissionId(c *gin.Context, submissionId int) string {
	s := submission.GetSubmissionsFromSubmissionId(c, submissionId)
	u := users.GetUserFromUserId(c, s.UserId)
	return u.UserName
}

func UserAuthorizator(data interface{}, c *gin.Context) bool {
	// 引数"data"はGetUserNameFromJwtのreturn
	if jwtUser, ok := data.(*myTypes.User); ok {
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
					urlUserName = GetUserNameFromsubmissionId(c, urlSubmissionId)
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
