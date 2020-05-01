package smtpclient

import (
	"bytes"
	"context"
	"net/smtp"
)

func New(smtpAddr string) (Client, func() error, error) {
	client, err := smtp.Dial(smtpAddr)
	if err != nil {
		return nil, nil, err
	}

	smtpClient := &smtpClient{api: client}
	return smtpClient, client.Quit, nil
}

type smtpClient struct {
	api *smtp.Client
}

func (c *smtpClient) Ping() error {
	return c.api.Noop()
}

func (c *smtpClient) SendMail(ctx context.Context, from, to, subject, msg string) error {
	err := c.api.Mail(from)
	err = c.resetIfError(err)
	if err != nil {
		return err
	}

	err = c.api.Rcpt(to)
	err = c.resetIfError(err)
	if err != nil {
		return err
	}

	wc, err := c.api.Data()
	err = c.resetIfError(err)
	if err != nil {
		return err
	}

	buf := bytes.NewBufferString(msg)
	_, err = buf.WriteTo(wc)
	if err != nil {
		return err
	}

	err = wc.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *smtpClient) resetIfError(err error) error {
	if err != nil {
		err2 := c.api.Reset()
		if err2 != nil {
			return err2
		}
		return err
	}
	return nil
}
