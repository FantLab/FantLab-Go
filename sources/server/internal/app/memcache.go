package app

import (
	"context"
	"fantlab/base/memcacheclient"
	"fmt"
	"strconv"
	"time"
)

const (
	userKey              = "users:user_id=%d"
	workStatKey          = "workstat%d"
	userResponseKey      = "user%dwork%dresp"
	homepageResponsesKey = "last:responses:home"
	articleLikeCountKey  = "blog:topic:likes:%d"
)

func (s *Services) DeleteUserCache(ctx context.Context, userId uint64) error {
	return s.deleteCache(ctx, func() string {
		return fmt.Sprintf(userKey, userId)
	})
}

func (s *Services) DeleteWorkStatCache(ctx context.Context, workId uint64) error {
	return s.deleteCache(ctx, func() string {
		return fmt.Sprintf(workStatKey, workId)
	})
}

func (s *Services) DeleteUserResponseCache(ctx context.Context, userId, workId uint64) error {
	return s.deleteCache(ctx, func() string {
		return fmt.Sprintf(userResponseKey, userId, workId)
	})
}

func (s *Services) DeleteHomepageResponsesCache(ctx context.Context) error {
	return s.deleteCache(ctx, func() string {
		return homepageResponsesKey
	})
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

func (s *Services) deleteCache(ctx context.Context, getKey func() string) error {
	if s.memcache == nil {
		return nil
	}
	err := s.memcache.Delete(ctx, getKey())
	if err != nil && !memcacheclient.IsNotFoundError(err) {
		return err
	}
	return nil
}
