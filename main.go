package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	sendMessage()
}

func sendMessage() {

	token := os.Getenv("BOT_TOKEN")
	chatid := os.Getenv("CHAT_ID")

	resp, err := http.Post("https://api.telegram.org/bot"+token+"/sendMessage?chat_id="+chatid+"&text=A push/pullrequest has been commited to main branch", "application/json", nil)

	if err != nil {
		log.Fatal(err, resp)
	}
}
