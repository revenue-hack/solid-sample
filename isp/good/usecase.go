package good

import "log"

type AdminAuthenticationUseCase struct {
	userInterface AdminUserInterface
}

func NewAdminAuthenticationUseCase(userI AdminUserInterface) *AdminAuthenticationUseCase {
	return &AdminAuthenticationUseCase{userInterface: userI}
}

func (use *AdminAuthenticationUseCase) do(name, password string) bool {
	u := &User{Name: name, Password: password}
	return use.userInterface.Authenticated(u)
}

type UserListUseCase struct {
	userInterface ClientUserInterface
}

func NewUserListUseCase(userI ClientUserInterface) *UserListUseCase {
	return &UserListUseCase{userInterface: userI}
}

func (use *UserListUseCase) do(name, password string) []*User {
	u := &User{Name: name, Password: password}
	// Admin専用の認証は使えない
	if use.userInterface.ClientAuthenticated(u) {
		return use.userInterface.List()
	}
	return nil
}

// Adminシステムでユーザ認証を行う
func adminAuthenticationRoute() {
	name := "test_name"
	password := "password"

	impl := &UserInterfaceImpl{}
	use := NewAdminAuthenticationUseCase(impl)
	use.do(name, password)
}

// クライアントシステムでユーザ認証を行い、ユーザ一覧を取得する
func userList() {
	name := "client_test_name"
	password := "password"

	impl := &UserInterfaceImpl{}
	use := NewUserListUseCase(impl)
	users := use.do(name, password)
	log.Println(users)
}
