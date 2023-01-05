package user

import "github.com/an1l4/testing-gomock-mockgen/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSometing(123, "Hello mock")
}
