package main

import (
	"context"
	"fmt"
	"log"

	pb "application/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "application/client/docs"
)

var c pb.MessageSenderClient

// @Summary send message 
// @ID send-message-ctrl
// @Produce json
// @Param data body 
// @Success 200 {object} send
// @Failure 400 {object} message
// @Router /send [post]

func SendMessageCtrl(ctx *gin.Context) {

	var newMessage pb.MessageRequest

	if err := ctx.ShouldBindJSON(&newMessage); err != nil {

		fmt.Println("Failed to accept message from cleint")
	}

	res, err := c.Sender(context.Background(), &pb.MessageRequest{
		Message:  newMessage.Message,
		Priority: newMessage.Priority,
	},)

	if err != nil {

		log.Fatalf("Failed to call Sender RPC: %v", err)
	}
	fmt.Println(res)
}

// @title Go + Gin Telegram Bot Api
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {

	fmt.Println("Client Side is working")
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to Dialing")
	}

	c = pb.NewMessageSenderClient(conn)
	fmt.Printf("Client is created", c) 

	router := gin.Default()

	router.POST("/send", SendMessageCtrl)
    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))	
    router.Run("localhost:8080")

}
