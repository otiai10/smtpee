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
	return &Client{Client: *c, To: []string{}}, err
}

// SetFrom ...
func (c *Client) SetFrom(from string) *Client {
	c.From = from
	return c
}

// AddTo ...
func (c *Client) AddTo(to ...string) *Client {
	c.To = append(c.To, to...)
	return c
}

// Send ...
func (c *Client) Send(body string) error {

	// From
	c.Mail(c.From)

	// To
	data := bytes.NewBuffer([]byte{})
	for _, to := range c.To {
		c.Rcpt(to)
		data.WriteString("To:" + to + rn)
	}

	// Subject
	data.WriteString("Subject:" + c.Subject + rn)

	// Body
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
