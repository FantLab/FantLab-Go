package memcacheclient

import (
	"context"
	"time"

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

func (aw *apmWrapper) Add(ctx context.Context, key string, value []byte, ttl time.Time) error {
	span, ctx := apm.StartSpan(ctx, "ADD "+key, aw.spanType)
	defer span.End()
	return aw.client.Add(ctx, key, value, ttl)
}

func (aw *apmWrapper) Set(ctx context.Context, key string, value []byte, ttl time.Time) error {
	span, ctx := apm.StartSpan(ctx, "SET "+key, aw.spanType)
	defer span.End()
	return aw.client.Set(ctx, key, value, ttl)
}

func (aw *apmWrapper) Get(ctx context.Context, key string) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "GET "+key, aw.spanType)
	defer span.End()
	return aw.client.Get(ctx, key)
}

func (aw *apmWrapper) Delete(ctx context.Context, key string) error {
	span, ctx := apm.StartSpan(ctx, "DELETE "+key, aw.spanType)
	defer span.End()
	return aw.client.Delete(ctx, key)
}
