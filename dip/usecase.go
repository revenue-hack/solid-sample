package dip

import "log"

type UserListUseCase struct {
	UserInterface UserInterface
}

func (use *UserListUseCase) Do(name, password string) []*User {
	u := &User{Name: name, Password: password}
	if use.UserInterface.Authenticated(u) {
		return use.UserInterface.List()
	}
	return nil
}

func userList() {
	name := "test_name"
	password := "password"

	impl := &UserInterfaceImpl{}
	use := &UserListUseCase{UserInterface: impl}
	users := use.Do(name, password)
	log.Println(users)
}
