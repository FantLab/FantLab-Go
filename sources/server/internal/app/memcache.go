package app

import (
	"context"
	"fantlab/base/memcacheclient"
	"fmt"
	"strconv"
	"time"
)

const (
	userKey             = "users:user_id=%d"
	articleLikeCountKey = "blog:topic:likes:%d"
)

func (s *Services) InvalidateUserCache(ctx context.Context, userId uint64) {
	if s.memcache == nil {
		return
	}
	_ = s.memcache.Delete(ctx, userCacheKeyPrefix+strconv.FormatUint(userId, 10))
}

func (s *Services) DeleteUserCache(ctx context.Context, userId uint64) error {
	if s.memcache == nil {
		return nil
	}
	err := s.memcache.Delete(ctx, fmt.Sprintf(userKey, userId))
	if err != nil && !memcacheclient.IsNotFoundError(err) {
		return err
	}
	return nil
}

func (s *Services) GetBlogArticleLikeCountCache(ctx context.Context, articleId uint64) (uint64, error) {
	if s.memcache == nil {
		return 0, nil
	}
	likeCountBytes, err := s.memcache.Get(ctx, fmt.Sprintf(articleLikeCountKey, articleId))
	if err != nil {
		return 0, err
	}
	likeCount, err := strconv.Atoi(string(likeCountBytes))
	return uint64(likeCount), err
}

func (s *Services) SetBlogArticleLikeCountCache(ctx context.Context, articleId, likeCount uint64) error {
	if s.memcache == nil {
		return nil
	}
	bytes := []byte(strconv.Itoa(int(likeCount)))
	return s.memcache.Set(ctx, fmt.Sprintf(articleLikeCountKey, articleId), bytes, time.Now().Add(1*time.Hour))
}
