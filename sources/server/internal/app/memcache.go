package app

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"
)

const articleLikeCountKey string = "blog:topic:likes:%d"

func (s *Services) GetBlogArticleLikeCount(ctx context.Context, articleId uint64) (uint64, error) {
	likeCountBytes, err := s.memcache.Get(ctx, fmt.Sprintf(articleLikeCountKey, articleId))

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(likeCountBytes), nil
}

func (s *Services) SetBlogArticleLikeCount(ctx context.Context, articleId, likeCount uint64) error {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, likeCount)
	return s.memcache.Set(ctx, fmt.Sprintf(articleLikeCountKey, articleId), bytes, time.Now().Add(1*time.Hour))
}
