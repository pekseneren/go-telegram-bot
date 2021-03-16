package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	repositoryName := os.Args[1]

	sendMessage(repositoryName)
}

func sendMessage(repositoryName string) {

	token := os.Getenv("BOT_TOKEN")
	chatid := os.Getenv("CHAT_ID")

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=A push/pullrequest has been commited to your (%s) repository on main branch", token, chatid, repositoryName)

	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		log.Fatal(err, resp)
	}
}
