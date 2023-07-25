package user

import "rzgonz.id/contract/user"

type UserUsecase interface {
	AddUser(user *model.User) (*model.User, error)
	FindUserById(id model.UserId) (*model.User, error)
	FindUsers() (*[]model.User, error)
	UpdateUser(user *model.UserUpdate) (*model.User, error)
	DeleteUser(id *model.UserId) error
}
