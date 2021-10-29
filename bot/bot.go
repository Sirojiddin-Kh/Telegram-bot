package bot

import (
	"log"
	"net/http"
	"bytes"
	"time"
	"encoding/json"
)

var chatUusername string = "@taskbotforudevs"
var telegramBotToken string = "2072163806:AAGw4j38aVS11P50aBBg8BIPoJiywmKZLHI"


type BotMessage struct {
	ChatUsername string `json:"chat_id"`
	Text   string `json:"text"`
}

func MessageSenderBot(message string) (err error) {

	var (
		addres string = "https://api.telegram.org/bot" + telegramBotToken
		text BotMessage 
	)

	text.ChatUsername = chatUusername
	text.Text = message

	buf, err := json.Marshal(text)

	if err != nil {

		log.Printf("Failed to Marshaling : %v", err)
		return  
	}
	time.Sleep(time.Second * 5)
	_, err = http.Post(addres + "/sendMessage", "application/json", bytes.NewBuffer(buf))

	if err != nil {

		log.Printf("Failed to send message : %v", err)
		return err

	}

	return 
}