package main

import (
	pb "application/proto"	
	bot "application/bot"
	"context"
	"fmt"
	"log"
	"net"	
	"time"
	_ "github.com/go-telegram-bot-api/telegram-bot-api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageSenderServer
}

var Messages []pb.MessageRequest

func (s *server) Sender(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {

	fmt.Println("Sender function was invoked")

	var message pb.MessageRequest

	message.Text = req.Text
	message.Priority = req.Priority

	Messages = append(Messages, message)

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
        for i, message := range Messages {
            if message.Priority == "high" {
               err := bot.MessageSenderBot(Messages[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    Messages = Remove(Messages, i)
                    time.Sleep(time.Second * 10) 
                    break
                } 
            }
        }
        
        if isSend {
            continue
        }

        for i, message := range Messages{
            if message.Priority == "medium" {
               err := bot.MessageSenderBot(Messages[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    Messages = Remove(Messages, i)
                    time.Sleep(time.Second * 10) 
                    break
                } 
            }
        }
        if isSend {
            continue
        }
        for i, message := range Messages{
        	if message.Priority == "low" {
                err := bot.MessageSenderBot(Messages[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    Messages = Remove(Messages, i)
                    time.Sleep(time.Second * 10) 
                    break
                } 
            }
        }


    }
}

func Remove(s []pb.MessageRequest, i int) []pb.MessageRequest {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}