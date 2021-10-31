package main

import (
	bot "application/bot"
	pb "application/proto"
	"context"
	"fmt"
	_ "github.com/go-telegram-bot-api/telegram-bot-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedMessageSenderServer
}

var HighPriorityMessages []pb.MessageRequest
var MediumPriorityMessages []pb.MessageRequest
var LowPriorityMessages []pb.MessageRequest

func (s *server) Sender(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {

	fmt.Println("Sender function was invoked")

	var message pb.MessageRequest

	message.Text = req.Text
	message.Priority = req.Priority

	if message.Priority == "high" {

		HighPriorityMessages = append(HighPriorityMessages, message)

	} else if message.Priority == "medium" {

		MediumPriorityMessages = append(MediumPriorityMessages, message)

	} else if message.Priority == "low" {

		LowPriorityMessages = append(LowPriorityMessages, message)
	} else {

		fmt.Println("Undefined Priority")
	}

	res := &pb.MessageResponse{

		Message: req.Text,
	}
	fmt.Println(message.Text)

	return res, nil
}

func main() {

	fmt.Println("Server Side is working")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {

		log.Fatalf("Failed to Listen the port : %v", err)
	}
	go Sender()

	s := grpc.NewServer()
	pb.RegisterMessageSenderServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Server : %v", err)
	}
}

func Sender() {
	var isSend bool
	for {
		isSend = false
		for i := range HighPriorityMessages {
			err := bot.MessageSenderBot(HighPriorityMessages[i].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			} else {
				isSend = true
				HighPriorityMessages = Remove(HighPriorityMessages, i)
				time.Sleep(time.Second * 10)
				break
			}
		}

		if isSend {
			continue
		}

		for i := range MediumPriorityMessages {
			err := bot.MessageSenderBot(MediumPriorityMessages[i].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			} else {
				isSend = true
				MediumPriorityMessages = Remove(MediumPriorityMessages, i)
				time.Sleep(time.Second * 10)
				break
			}

		}
		if isSend {
			continue
		}
		for i := range LowPriorityMessages {
			err := bot.MessageSenderBot(LowPriorityMessages[i].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			} else {
				isSend = true
				LowPriorityMessages = Remove(LowPriorityMessages, i)
				time.Sleep(time.Second * 10)
				break
			}
		}

	}
}

func Remove(s []pb.MessageRequest, i int) []pb.MessageRequest {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
