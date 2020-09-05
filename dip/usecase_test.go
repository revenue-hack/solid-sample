package dip_test

import (
	"reflect"
	"testing"

	"github.com/revenue-hack/solid-sample/dip"
)

var (
	expectedUsers = []*dip.User{
		{
			Name:     "user1",
			Password: "password1",
		},
		{
			Name:     "user2",
			Password: "password2",
		},
	}
)

func TestBad_UserList(t *testing.T) {
	name := "test_name"
	password := "password"

	impl := &dip.UserInterfaceImpl{}
	use := &dip.UserListUseCase{UserInterface: impl}
	users := use.Do(name, password)

	expected := []*dip.User{
		{
			Name:     "user1",
			Password: "password1",
		},
		{
			Name:     "user2",
			Password: "password2",
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("users are not match. result is %v, expected is %v", users, expected)
	}
}

func TestGood_UserList(t *testing.T) {
	name := "test_name"
	password := "password"

	use := &dip.UserListUseCase{UserInterface: &UserInterfaceMockImpl{}}
	users := use.Do(name, password)

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("users are not match. result is %v, expected is %v", users, expectedUsers)
	}
}

type UserInterfaceMockImpl struct{}

func (impl *UserInterfaceMockImpl) Authenticated(u *dip.User) bool {
	return true
}

func (impl *UserInterfaceMockImpl) List() []*dip.User {
	return expectedUsers
}
