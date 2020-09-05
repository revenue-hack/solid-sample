package bad

import "log"

type AdminAuthenticationUseCase struct {
	userInterface UserInterface
}

func (use *AdminAuthenticationUseCase) do(name, password string) bool {
	u := &User{Name: name, Password: password}
	return use.userInterface.Authenticated(u)
}

type UserListUseCase struct {
	userInterface UserInterface
}

func (use *UserListUseCase) do(name, password string) []*User {
	u := &User{Name: name, Password: password}
	// 実はこの認証はAdmin専用でこのユーザ認証は間違っている
	if use.userInterface.Authenticated(u) {
		return use.userInterface.List()
	}
	return nil
}

// Adminシステムでユーザ認証を行う
func adminAuthenticationRoute() {
	name := "test_name"
	password := "password"

	impl := UserInterfaceImpl{}
	use := &AdminAuthenticationUseCase{}
	use.do(name, password)
}

// クライアントシステムでユーザ認証を行い、ユーザ一覧を取得する
func userList() {
	name := "test_name"
	password := "password"

	impl := UserInterfaceImpl{}
	use := &UserListUseCase{}
	users := use.do(name, password)
	log.Println(users)
}
