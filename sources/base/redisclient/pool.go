package redisclient

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewPool(server string, maxIdle uint8) (Client, func() error) {
	pool := &redis.Pool{
		MaxIdle: int(maxIdle),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", server)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &client{pool: pool}, pool.Close
}

type client struct {
	pool *redis.Pool
}

func (c *client) Perform(ctx context.Context, fn func(conn Conn) error) error {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(conn)
}
