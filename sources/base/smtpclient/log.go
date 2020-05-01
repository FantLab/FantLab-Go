package smtpclient

import "context"

type LogEntry struct {
	To      string
	Subject string
	Err     error
}

type LogFunc func(context.Context, *LogEntry)

func Log(client Client, fn LogFunc) Client {
	if client == nil {
		return nil
	}
	return &logWrapper{
		client: client,
		fn:     fn,
	}
}

type logWrapper struct {
	client Client
	fn     LogFunc
}

func (lw *logWrapper) Ping() error {
	return lw.client.Ping()
}

func (lw *logWrapper) SendMail(ctx context.Context, from, to, subject, msg string) error {
	err := lw.client.SendMail(ctx, from, to, subject, msg)
	lw.fn(ctx, &LogEntry{
		To:      to,
		Subject: subject,
		Err:     err,
	})
	return err
}
