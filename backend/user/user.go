package user

import "time"

type User struct {
	UserId      int64
	Email       string
	Password    string
	CreatedDate time.Time
	Role        string
}
