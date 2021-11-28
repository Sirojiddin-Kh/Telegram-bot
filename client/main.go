package main

import (
	"context"
	"fmt"
	"log"
	_ "application/client/docs"
	"github.com/gin-gonic/gin"
	pb "application/proto"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

var c pb.MessageSenderClient

// @Summary MessageSender
// @Description send message
// @ID message_sender
// @Accept json
// @Produce json
// @Param input body Message true "message_content"
// @Success 200
// @Failure 400
// @Router /send [post]


func SendMessageHandler(ctx *gin.Context) {

	var newMessage pb.MessageRequest

	if err := ctx.BindJSON(&newMessage); err != nil {

		log.Printf("Failed to accept message from cleint")
	}
	fmt.Println(newMessage)
	res, err := c.Sender(context.Background(), &pb.MessageRequest{
		Text:  newMessage.Text,
		Priority: newMessage.Priority,
	})

	if err != nil {

		log.Fatalf("Failed to call Sender RPC: %v", err)
	}
	log.Println(res)
}


// @title Message Sender Bot
// @version 1.0
// @description Telegram Bot which sends messages to channels and groups
// @host localhost:8000
// @BasePath /

func main() {

	fmt.Println("Client Side is working")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to Dialing")
	}

	c = pb.NewMessageSenderClient(conn)
	fmt.Printf("Client is created", c)

	router := gin.Default()

	router.POST("/send", SendMessageHandler)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")

}
