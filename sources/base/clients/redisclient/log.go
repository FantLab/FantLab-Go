package redisclient

import "context"

type LogFunc func(context.Context, error)

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

func (lw *logWrapper) Perform(ctx context.Context, fn func(conn Conn) error) error {
	err := lw.client.Perform(ctx, fn)
	if err != nil {
		lw.fn(ctx, err)
	}
	return err
}
