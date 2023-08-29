package user

import (
	userContract "github.com/rzgonz/samplego/contract/user"
)

type UserUsecase interface {
	AddUser(user *userContract.User) (*userContract.User, error)
	FindUserById(id userContract.UserId) (*userContract.User, error)
	FindUsers() (*[]userContract.User, error)
	UpdateUser(user *userContract.UserUpdate) (*userContract.User, error)
	DeleteUser(id *userContract.UserId) error
}
