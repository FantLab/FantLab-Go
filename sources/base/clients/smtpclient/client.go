package smtpclient

import (
	"bytes"
	"context"
	"github.com/FantLab/go-kit/codeflow"
	"io"
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
	var wc io.WriteCloser

	err := codeflow.Try(
		func() error {
			return c.api.Mail(from)
		},
		func() error {
			return c.api.Rcpt(to)
		},
		func() error {
			w, err := c.api.Data()
			wc = w
			return err
		},
		func() error {
			buf := bytes.NewBufferString(msg)
			_, err := buf.WriteTo(wc)
			return err
		},
		func() error {
			return wc.Close()
		},
	)

	if err != nil {
		_ = c.api.Reset()
		return err
	}

	return nil
}
