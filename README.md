# smtpee
smtpeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee

```go
client, _ := smtpee.Dial("localhost:25")

client.From = "otiai10+sender@gmail.com"
client.To = "otiai10+recipient@gmail.com"
client.Subject = "Hi"

client.Send("I'm hungry.")
```
