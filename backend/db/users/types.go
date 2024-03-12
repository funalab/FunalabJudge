package users

import (
	"go-test/db/problems"
	"time"
)

type User struct {
	UserName   string    `bson:"userName"`
	Password   string    `bson:"password"`
	JoinedDate time.Time `bson:"joinedDate"`
}
type UserStatus struct {
	UserName       string                       `bson:"userName"`
	ProblemsStatus []problems.ProblemWithStatus `bson:"problemsStatus"`
}
