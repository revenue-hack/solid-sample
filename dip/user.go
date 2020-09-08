package dip

import "errors"

type User struct {
	name     string
	password string
}

func (u *User) Name() string {
	return u.name
}
func (u *User) Password() string {
	return u.password
}

type UserInterface interface {
	Authenticated(u *User) bool
	List() []*User
}

func NewUser(name, password string) (*User, error) {
	// 名前は必須、10文字以上15文字以内
	// パスワードは必須, 12文字以内
	if name == "" {
		return nil, errors.New("名前が不正です")
	}
	if len(name) < 10 || len(name) > 15 {
		return nil, errors.New("名前が不正です")
	}
	if password == "" {
		return nil, errors.New("パスワードが不正です")
	}
	if len(password) > 12 {
		return nil, errors.New("パスワードがが不正です")
	}

	return &User{name: name, password: password}, nil
}

type UserInterfaceImpl struct{}

func (impl *UserInterfaceImpl) Authenticated(u *User) bool {
	return u.Name() == "test_name" && u.Password() == "password"
}

func (impl *UserInterfaceImpl) List() []*User {
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
