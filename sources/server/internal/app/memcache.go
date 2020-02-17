package app

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

const articleLikeCountKey string = "blog:topic:likes:%d"

func (s *Services) GetBlogArticleLikeCount(ctx context.Context, articleId uint64) (uint64, error) {
	likeCountBytes, err := s.memcache.Get(ctx, fmt.Sprintf(articleLikeCountKey, articleId))

	if err != nil {
		return 0, err
	}

	likeCount, err := strconv.Atoi(string(likeCountBytes))

	return uint64(likeCount), err
}

func (s *Services) SetBlogArticleLikeCount(ctx context.Context, articleId, likeCount uint64) error {
	bytes := []byte(strconv.Itoa(int(likeCount)))
	return s.memcache.Set(ctx, fmt.Sprintf(articleLikeCountKey, articleId), bytes, time.Now().Add(1*time.Hour))
}
