package users

import "time"
import "go-test/db/submission"

type User struct {
	UserName   string    `bson:"userName"`
	Password   string    `bson:"password"`
	JoinedDate time.Time `bson:"joinedDate"`
}
type UserStatus struct {
	UserName       string              `bson:"userName"`
	ProblemsStatus []submission.Result `bson:"problemsStatus"`
}
