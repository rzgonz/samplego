package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rzgonz/samplego/contract/user"
	"github.com/rzgonz/samplego/rzuser/domain/usecase"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func CreateUserHandler(gr *grpc.Server, userUsecase usecase.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	user.RegisterUsersServer(gr, &userHandler)
}

func (e *UserHandler) GetUserList(ctx context.Context, in *empty.Empty) (*user.UserList, error) {
	users, err := e.userUsecase.FindUsers()
	if err != nil {
		return nil, err
	}
	var userx = make([]*user.User, 0)
	for i := 0; i < len(*users); i++ {
		var data = new(user.User)
		data.Id = (*users)[i].Id
		data.Email = (*users)[i].Email
		data.Name = (*users)[i].Name
		data.Alamat = (*users)[i].Alamat
		data.Password = (*users)[i].Password
		userx = append(userx, data)
	}
	var u = user.UserList{
		List: userx,
	}
	return &u, nil
}
func (e *UserHandler) GetUserById(ctx context.Context, in *user.UserId) (*user.User, error) {
	user, err := e.userUsecase.FindUserById(*in)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (e *UserHandler) InsertUser(ctx context.Context, in *user.User) (*empty.Empty, error) {
	_, err := e.userUsecase.AddUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
func (e *UserHandler) UpdateUser(ctx context.Context, in *user.UserUpdate) (*empty.Empty, error) {
	_, err := e.userUsecase.UpdateUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
func (e *UserHandler) DeleteUser(ctx context.Context, in *user.UserId) (*empty.Empty, error) {
	err := e.userUsecase.DeleteUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}

// Implement the mustEmbedUnimplementedUsersServer method.
func (e *UserHandler) UnimplementedUsersServer() {
	panic("error")
}
