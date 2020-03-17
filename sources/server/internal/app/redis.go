package app

import (
	"context"
	"fantlab/base/redisco"
	"fmt"
)

func (s *Services) BlogTopicsViewCount(ctx context.Context, topicIds []uint64) []uint64 {
	if s.redis == nil {
		return nil
	}
	result := make([]uint64, len(topicIds))
	_ = s.redis.Perform(ctx, func(conn redisco.Conn) error {
		for index, topicId := range topicIds {
			viewCount, _ := redisco.Uint64(conn.Do("PFCOUNT", fmt.Sprintf("blogtopicviews:%d", topicId)))
			result[index] = viewCount
		}
		return nil
	})
	return result
}
