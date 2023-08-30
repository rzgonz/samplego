package main

import (
	"log"

	"github.com/rzgonz/samplego/contract/user"
	userHandler "github.com/rzgonz/samplego/product/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	port := "8827"
	targetPort := "8957"
	conn, err := grpc.Dial(":"+targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cloud not conect to %v %v", targetPort, err)
	}
	user := user.NewUsersClient(conn)
	router := gin.Default()

	userHandler.CreateUserHandler(router, user)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
