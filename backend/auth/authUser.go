package auth

import (
	"fmt"
	"go-test/submission"
	"go-test/types"
	"go-test/user"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserNameFromJwt(c *gin.Context) interface{} {
	// jwtからuserNameを抽出し、UserAuthorizatorに渡す
	claims := jwt.ExtractClaims(c)
	return &types.User{
		UserName: claims[jwt.IdentityKey].(string),
	}
}

func GetUserNameFromsubmissionId(c *gin.Context, submissionId int) string {
	s := submission.GetSubmissionsFromSubmissionId(c, submissionId)
	u := user.GetUserFromUserId(c, s.UserId)
	return u.UserName
}

func UserAuthorizator(data interface{}, c *gin.Context) bool {
	// 引数"data"はGetUserNameFromJwtのreturn
	if v, ok := data.(*types.User); ok {
		jwtUserName := v.UserName
		jwtUser := user.GetUserFromUserName(c, jwtUserName)
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
			if jwtUserName == urlUserName {
				return true
			}
		} else {
			// unexpected Role
			return false
		}
	}

	return false
}
