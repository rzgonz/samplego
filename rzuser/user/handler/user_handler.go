package handler

import (
	"context"
	"rzuser/user"

	"github.com/rzgonz/samplego/contract/user/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userUsecase user.UserUsecase
	// Embed the unimplemented server

}

func CreateUserHandler(gr *grpc.Server, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	model.RegisterUsersServer(gr, &userHandler)
}

func (e *UserHandler) GetUserList(ctx context.Context, in *empty.Empty) (*model.UserList, error) {
	users, err := e.userUsecase.FindUsers()
	if err != nil {
		return nil, err
	}
	var userx = make([]*model.User, 0)
	for i := 0; i < len(*users); i++ {
		var data = new(model.User)
		data.Id = (*users)[i].Id
		data.Email = (*users)[i].Email
		data.Name = (*users)[i].Name
		data.Alamat = (*users)[i].Alamat
		data.Password = (*users)[i].Password
		userx = append(userx, data)
	}
	var u = model.UserList{
		List: userx,
	}
	return &u, nil
}
func (e *UserHandler) GetUserById(ctx context.Context, in *model.UserId) (*model.User, error) {
	user, err := e.userUsecase.FindUserById(*in)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (e *UserHandler) InsertUser(ctx context.Context, in *model.User) (*empty.Empty, error) {
	_, err := e.userUsecase.AddUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
func (e *UserHandler) UpdateUser(ctx context.Context, in *model.UserUpdate) (*empty.Empty, error) {
	_, err := e.userUsecase.UpdateUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
func (e *UserHandler) DeleteUser(ctx context.Context, in *model.UserId) (*empty.Empty, error) {
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
