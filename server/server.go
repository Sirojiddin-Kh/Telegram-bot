package main

import (
	pb "application/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	_ "github.com/go-telegram-bot-api/telegram-bot-api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageSenderServer
}

type BotMessage struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func (s *server) Sender(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {

	fmt.Println("Sender function was invoked")

	chat_id := "@taskbotforudevs"

	res := &pb.MessageResponse{

		ChatId:  chat_id,
		Message: req.Message,
	}

	var botMessage BotMessage
	addres := "https://api.telegram.org/bot" + "2072163806:AAGw4j38aVS11P50aBBg8BIPoJiywmKZLHI"

	botMessage.ChatId = "@taskbotforudevs"
	botMessage.Text = req.Message

	buf, err := json.Marshal(botMessage)

	if err != nil {
		log.Fatalf("Failed in Json: %v", err)
	}

	if req.Priority == "high" {

		time.Sleep(5 * time.Second)
		_, err = http.Post(addres+"/sendMessage", "application/json", bytes.NewBuffer(buf))

	} else if req.Priority == "medium" {

		time.Sleep(7 * time.Second)
		_, err = http.Post(addres+"/sendMessage", "application/json", bytes.NewBuffer(buf))

	} else if req.Priority == "low" {

		time.Sleep(10 * time.Second)
		_, err = http.Post(addres+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	} else {

		fmt.Println("Undefined priority")
	}

	if err != nil {
		log.Fatalf("Failed in Sending: %v", err)
	}

	return res, nil
}

func main() {

	fmt.Println("Server Side is working")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {

		log.Fatalf("Failed to Listen the port : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageSenderServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Server : %v", err)
	}
}
