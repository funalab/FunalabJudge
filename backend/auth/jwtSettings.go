package auth

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
)

const JwtIdentityKey = "user"
const JwtJoinedDateKey = "joined"

func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "funalabJudge",
		Key:             []byte(os.Getenv("SECRET_KEY")), // 運用時には再作成する: % openssl rand -base64 32
		Timeout:         time.Hour * 24 * 7,              // equals to CookieMaxAge
		MaxRefresh:      time.Hour * 24 * 7,
		SendCookie:      true,
		SecureCookie:    false, //non HTTPS dev environments
		CookieHTTPOnly:  true,  // JS can't modify
		CookieDomain:    os.Getenv("SERVER_IP"),
		CookieName:      "token", // default:jwt
		TokenLookup:     "cookie:token",
		CookieSameSite:  http.SameSiteStrictMode, // CSRF対策にもなる
		IdentityKey:     JwtIdentityKey,
		PayloadFunc:     JwtMapper,
		Authenticator:   LoginAuthenticator,
		IdentityHandler: GetUserNameFromJwt,
		Authorizator:    UserAuthorizator,
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
