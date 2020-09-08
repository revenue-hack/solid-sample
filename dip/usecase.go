package dip

import (
	"log"
)

type UserListUseCase struct {
	UserInterface UserInterface
}

func (use *UserListUseCase) Do(name, password string) []*User {
	u, err := NewUser(name, password)
	if err != nil {
		return nil
	}

	if use.UserInterface.Authenticated(u) {
		return use.UserInterface.List()
	}
	return nil
}

func userList() {
	name := "test_name"
	password := "password"

	impl := &UserInterfaceImplMock{}
	use := &UserListUseCase{UserInterface: impl}
	users := use.Do(name, password)

	log.Println(users)
}

type UserInterfaceImplMock struct{}

func (impl *UserInterfaceImplMock) Authenticated(u *User) bool {
	return true
}

func (impl *UserInterfaceImplMock) List() []*User {
	return []*User{
		{
			name:     "user1",
			password: "password1",
		},
		{
			name:     "user2",
			password: "password2",
		},
	}
}
