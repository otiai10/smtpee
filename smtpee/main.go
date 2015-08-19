package main

import (
	"log"

	"github.com/otiai10/smtpee"
)

func main() {

	c := smtpee.Client{
		Addr: "localhost:25",
		From: "otiai10@otiai10.com",
		To: []string{
			"otiai10@gmail.com",
		},
		Subject: "MacのSMTPのテスト",
	}

	err := c.Send("はらへった\nなんか食わせろ")

	log.Println(err)
}
