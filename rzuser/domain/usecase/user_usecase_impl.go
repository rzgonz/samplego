package usecase

import (
	"github.com/rzgonz/samplego/contract/user"
	"github.com/rzgonz/samplego/rzuser/domain"
)

type UserUsecaseImpl struct {
	userRepo domain.UserRepo
}

func CreateUserUsecase(userRepo domain.UserRepo) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) AddUser(user *user.User) (*user.User, error) {
	return e.userRepo.AddUser(user)
}

func (e *UserUsecaseImpl) FindUserById(id user.UserId) (*user.User, error) {
	return e.userRepo.FindUserById(id)
}

func (e *UserUsecaseImpl) FindUsers() (*[]user.User, error) {
	return e.userRepo.FindUsers()
}

func (e *UserUsecaseImpl) UpdateUser(user *user.UserUpdate) (*user.User, error) {
	return e.userRepo.UpdateUser(user)
}

func (e *UserUsecaseImpl) DeleteUser(id *user.UserId) error {
	return e.userRepo.DeleteUser(id)
}
