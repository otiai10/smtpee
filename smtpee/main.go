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

	c.SetFrom("Hiromu <hiromu+sender@origami.co>")
	c.AddTo("otiai10+recipient@gmail.com", "hiromu@origami.co")
	c.Subject = "元気ですか？"

	c.Send("はらへった\nなんか食わせろ")
}
