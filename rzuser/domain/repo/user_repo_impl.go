package repo

import (
	"github.com/rzgonz/samplego/rzuser/presistance/model"

	"github.com/rzgonz/samplego/contract/user"
	"github.com/rzgonz/samplego/rzuser/domain"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func CreateUserRepoImpl(db *gorm.DB) domain.UserRepo {
	return &UserRepoImpl{db}
}

func (e *UserRepoImpl) AddUser(user *user.User) (*user.User, error) {
	var userDB = model.UserDB{
		Name:     user.Name,
		Email:    user.Email,
		Alamat:   user.Alamat,
		Password: user.Password,
	}
	err := e.db.Table("user").Save(&userDB).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *UserRepoImpl) FindUserById(id user.UserId) (*user.User, error) {
	var user user.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (e *UserRepoImpl) FindUsers() (*[]user.User, error) {
	var users []user.User
	err := e.db.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (e *UserRepoImpl) UpdateUser(userData *user.UserUpdate) (*user.User, error) {
	var us user.User
	err := e.db.Table("user").Where("id = ?", userData.Id).First(&us).Error
	if err != nil {
		return nil, err
	}

	// Update the user fields with the values from user.User
	err = e.db.Table("user").Model(&us).Updates(&userData.User).Error
	if err != nil {
		return nil, err
	}

	return &us, nil
}

func (e *UserRepoImpl) DeleteUser(id *user.UserId) error {
	var user user.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
