package users

import "time"

type User struct {
	UserId      int32     `bson:"userId"`
	UserName    string    `bson:"userName"`
	Password    string    `bson:"password"`
	CreatedDate time.Time `bson:"createdDate"`
	Role        string    `bson:"role"`
}
