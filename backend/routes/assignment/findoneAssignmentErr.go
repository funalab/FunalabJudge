package assignment

import "fmt"

type FindOneAssignmentErr struct {
	msg string
}

func (f *FindOneAssignmentErr) Error() string {
	return fmt.Sprintf("FindOneAssignmentErr: %v\n", f.msg)
}

func NewFindOneAssignmentErr(errMsg string) error {
	return &FindOneAssignmentErr{
		msg: errMsg,
	}
}
