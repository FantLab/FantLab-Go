package smtpclient

import (
	"context"
	"fmt"
	"go.elastic.co/apm"
)

func APM(client Client, spanType string) Client {
	if client == nil {
		return nil
	}
	return &apmWrapper{
		client:   client,
		spanType: spanType,
	}
}

type apmWrapper struct {
	client   Client
	spanType string
}

func (aw *apmWrapper) Ping() error {
	return aw.client.Ping()
}

func (aw *apmWrapper) SendMail(ctx context.Context, from, to, subject, msg string) error {
	span, ctx := apm.StartSpan(ctx, fmt.Sprintf("SEND MAIL; SUBJECT: %s, TO: %s", subject, to), aw.spanType)
	defer span.End()
	return aw.client.SendMail(ctx, from, to, subject, msg)
}
