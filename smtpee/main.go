package main

import (
	"log"

	"github.com/otiai10/smtpee"
)

func main() {
	c, err := smtpee.Dial("localhost:25")
	if err != nil {
		log.Fatalln(err)
	}

	c.From = "otiai10+sender@origami.co"
	c.To = []string{"otiai10+recipient@gmail.com"}
	c.Subject = "元気ですか？"

	c.Send("はらへった\nなんか食わせろ")
}
