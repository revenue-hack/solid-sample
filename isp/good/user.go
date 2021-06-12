package good

type User struct {
	Name     string
	Password string
}

type AdminUserInterface interface {
	Authenticated(u *User) bool
}
type ClientUserInterface interface {
	ClientAuthenticated(u *User) bool
	List() []*User
}

type UserInterfaceImpl struct{}

func (impl *UserInterfaceImpl) Authenticated(u *User) bool {
	return u.Name == "test_name" && u.Password == "password"
}
func (impl *UserInterfaceImpl) ClientAuthenticated(u *User) bool {
	return u.Name == "client_test_name" && u.Password == "password"
}

func (impl *UserInterfaceImpl) List() []*User {
	return []*User{
		{
			Name:     "user1",
			Password: "password1",
		},
		{
			Name:     "user2",
			Password: "password2",
		},
	}
}
