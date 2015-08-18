package smtpee

import (
	"bytes"
	"net/smtp"
)

const (
	rn = "\r\n"
)

// Client ...
type Client struct {
	smtp.Client
	From    string
	To      []string
	Subject string
}

// Dial ...
func Dial(addr string) (*Client, error) {
	c, err := smtp.Dial(addr)
	if err != nil {
		return nil, err
	}
	return &Client{Client: *c}, err
}

// Send ...
func (c *Client) Send(body string) error {

	data := bytes.NewBuffer([]byte{})
	c.Mail(c.From) // メールの送り主を指定
	data.WriteString("To:" + c.From + rn)
	for _, to := range c.To {
		c.Rcpt(to) // Ccにしたい場合も同様に指定
		data.WriteString("Cc:" + to + rn)
	}

	data.WriteString("Subject:" + c.Subject + rn)

	data.WriteString(body)

	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	if _, err = data.WriteTo(wc); err != nil {
		return err
	}

	return c.Quit()
}
