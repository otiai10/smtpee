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
	Addr    string
	From    string
	To      []string
	Subject string
}

// Send ...
func (c *Client) Send(body string) error {

	// smtp client
	sc, err := smtp.Dial(c.Addr)
	if err != nil {
		return err
	}

	// From
	sc.Mail(c.From)

	// To
	data := bytes.NewBuffer([]byte{})
	for _, to := range c.To {
		sc.Rcpt(to)
		data.WriteString("To:" + to + rn)
	}

	// Subject
	data.WriteString("Subject:" + c.Subject + rn)

	// Body
	data.WriteString(body)

	w, err := sc.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err = data.WriteTo(w); err != nil {
		return err
	}

	return sc.Quit()
}
