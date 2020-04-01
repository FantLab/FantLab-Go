package app

import (
	"context"
	"fantlab/base/redisclient"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func (s *Services) BlogTopicsViewCount(ctx context.Context, topicIds []uint64) []uint64 {
	result := make([]uint64, len(topicIds))
	if s.redis == nil {
		return result
	}
	_ = s.redis.Perform(ctx, func(conn redisclient.Conn) error {
		for index, topicId := range topicIds {
			viewCount, _ := redis.Uint64(conn.Do("PFCOUNT", fmt.Sprintf("blogtopicviews:%d", topicId)))
			result[index] = viewCount
		}
		return nil
	})
	return result
}
