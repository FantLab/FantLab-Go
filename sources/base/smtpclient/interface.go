package smtpclient

import "context"

type Client interface {
	Ping() error
	SendMail(ctx context.Context, from, to, subject, msg string) error
}
