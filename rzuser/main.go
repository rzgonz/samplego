package main

import (
	"fmt"
	"log"
	"net"
	userHandler "rzuser/user/handler"
	userRepo "rzuser/user/repo"
	userUsecase "rzuser/user/usecase"

	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"

	userDbModel "rzuser/model"

	"google.golang.org/grpc"
)

func main() {
	var port = "8957"
	dsn := "rzgonz:enter@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(
		userDbModel.UserDB{},
	)

	server := grpc.NewServer()

	userRepo := userRepo.CreateUserRepoImpl(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)

	userHandler.CreateUserHandler(server, userUsecase)

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Starting at Port: ", port)
	log.Fatal(server.Serve(conn))
}
