package myTypes

import "fmt"

/*MongoConnectionErr*/
type MongoConnectionErr struct {
	Msg string
}

func (c *MongoConnectionErr) Error() string {
	return fmt.Sprintf("ConnectionErr: %v\n", c.Msg)
}

func NewMongoConnectionErr(errMsg string) error {
	return &MongoConnectionErr{
		Msg: errMsg,
	}
}

/*FindRootDirErr*/
type FindRootDirErr struct {
	Msg string
}

func (f *FindRootDirErr) Error() string {
	return fmt.Sprintf("ConnectionErr: %v\n", f.Msg)
}

func NewFindRootDirErr(errMsg string) error {
	return &FindRootDirErr{
		Msg: errMsg,
	}
}

/*GetCurrentDirErr*/
type GetCurrentDirErr struct {
	Msg string
}

func (g *GetCurrentDirErr) Error() string {
	return fmt.Sprintf("ConnectionErr: %v\n", g.Msg)
}

func NewGetCurrentDirErr(errMsg string) error {
	return &GetCurrentDirErr{
		Msg: errMsg,
	}
}

/*FindOneAssignmentErr*/
type FindOneAssignmentErr struct {
	Msg string
}

func (f *FindOneAssignmentErr) Error() string {
	return fmt.Sprintf("FindOneAssignmentErr: %v\n", f.Msg)
}

func NewFindOneAssignmentErr(errMsg string) error {
	return &FindOneAssignmentErr{
		Msg: errMsg,
	}
}

/*MakeFileErr*/
type GenerateMakefileErr struct {
	Msg string
}

func (g *GenerateMakefileErr) Error() string {
	return fmt.Sprintf("Generate Makefile Error: %v\n", g.Error())
}

func NewGenerateMakefileErr(errMsg string) error {
	return &GenerateMakefileErr{
		Msg: errMsg,
	}
}

/* MakeFailErr*/
type MakeFailErr struct {
	Msg string
}

func (m *MakeFailErr) Error() string {
	return fmt.Sprintf("Make Fail Error: %v\n", m.Msg)
}

func NewMakeFailErr(errMsg string) error {
	return &MakeFailErr{
		Msg: errMsg,
	}
}
