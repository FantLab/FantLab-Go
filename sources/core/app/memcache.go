package app

import (
	"context"
	"fantlab/base/clients/memcacheclient"
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

var noExpiration = time.Unix(0, 0)

func (s *Services) DeleteUserCache(ctx context.Context, userId uint64) error {
	return memcacheclient.Perform(s.memcache, true, func(client memcacheclient.Client) error {
		return client.Delete(ctx, fmt.Sprintf(userKey, userId))
	})
}

func (s *Services) DeleteWorkStatCache(ctx context.Context, workId uint64) error {
	return memcacheclient.Perform(s.memcache, true, func(client memcacheclient.Client) error {
		return client.Delete(ctx, fmt.Sprintf(workStatKey, workId))
	})
}

func (s *Services) SetUserResponseCache(ctx context.Context, userId, workId uint64) error {
	return memcacheclient.Perform(s.memcache, false, func(client memcacheclient.Client) error {
		bytes := []byte(strconv.Itoa(1))
		return s.memcache.Set(ctx, fmt.Sprintf(userResponseKey, userId, workId), bytes, noExpiration)
	})
}

func (s *Services) DeleteUserResponseCache(ctx context.Context, userId, workId uint64) error {
	return memcacheclient.Perform(s.memcache, true, func(client memcacheclient.Client) error {
		return client.Delete(ctx, fmt.Sprintf(userResponseKey, userId, workId))
	})
}

func (s *Services) DeleteHomepageResponsesCache(ctx context.Context) error {
	return memcacheclient.Perform(s.memcache, true, func(client memcacheclient.Client) error {
		return client.Delete(ctx, homepageResponsesKey)
	})
}

func (s *Services) GetBlogArticleLikeCountCache(ctx context.Context, articleId uint64) (count uint64, err error) {
	err = memcacheclient.Perform(s.memcache, true, func(client memcacheclient.Client) error {
		b, e := s.memcache.Get(ctx, fmt.Sprintf(articleLikeCountKey, articleId))
		if e == nil {
			count, e = strconv.ParseUint(string(b), 10, 0)
		}
		return e
	})
	return
}

func (s *Services) SetBlogArticleLikeCountCache(ctx context.Context, articleId, likeCount uint64) error {
	return memcacheclient.Perform(s.memcache, false, func(client memcacheclient.Client) error {
		bytes := []byte(strconv.Itoa(int(likeCount)))
		return s.memcache.Set(ctx, fmt.Sprintf(articleLikeCountKey, articleId), bytes, time.Now().Add(1*time.Hour))
	})
}
