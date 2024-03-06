package users

import "time"

type User struct {
	UserName   string    `bson:"userName"`
	Password   string    `bson:"password"`
	JoinedDate time.Time `bson:"joinedDate"`
}
