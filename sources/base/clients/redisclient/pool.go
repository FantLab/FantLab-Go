package redisclient

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

type GetConnFunc = func(*redis.Pool, context.Context) (redis.Conn, error)

func NewPoolClient(server string, maxIdle uint8, getConn GetConnFunc) (Client, func() error) {
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
	return &client{pool: pool, getConn: getConn}, pool.Close
}

type client struct {
	pool    *redis.Pool
	getConn GetConnFunc
}

func (c *client) Perform(ctx context.Context, fn func(conn Conn) error) error {
	conn, err := c.getConn(c.pool, ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(conn)
}
