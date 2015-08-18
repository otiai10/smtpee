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

	c.SetFrom("otiai10+sender@gmail.com")
	c.AddTo("otiai10+recipient02@gmail.com", "otiai10+recipient01@gmail.com")
	c.Subject = "元気ですか？"

	c.Send("はらへった\nなんか食わせろ")
}
