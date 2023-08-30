package usecase

import (
	"github.com/rzgonz/samplego/rzuser/model"
	"github.com/rzgonz/samplego/rzuser/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) AddUser(user *model.User) (*model.User, error) {
	return e.userRepo.AddUser(user)
}

func (e *UserUsecaseImpl) FindUserById(id model.UserId) (*model.User, error) {
	return e.userRepo.FindUserById(id)
}

func (e *UserUsecaseImpl) FindUsers() (*[]model.User, error) {
	return e.userRepo.FindUsers()
}

func (e *UserUsecaseImpl) UpdateUser(user *model.UserUpdate) (*model.User, error) {
	return e.userRepo.UpdateUser(user)
}

func (e *UserUsecaseImpl) DeleteUser(id *model.UserId) error {
	return e.userRepo.DeleteUser(id)
}
