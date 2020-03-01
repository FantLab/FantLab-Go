package redisco

import "context"

type Conn interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
	Send(commandName string, args ...interface{}) error
	Flush() error
	Receive() (reply interface{}, err error)
}

type Client interface {
	Perform(ctx context.Context, fn func(conn Conn) error) error
}
