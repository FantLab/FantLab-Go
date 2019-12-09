package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"testing"
	"time"
)

func Test_LikeDislike(t *testing.T) {
	stubDB := &dbstubs.StubDB{
		ExecTable:  make(dbstubs.StubExecTable),
		QueryTable: make(dbstubs.StubQueryTable),
	}

	db := NewDB(stubDB)

	t.Run("like_dislike", func(t *testing.T) {
		{
			stubDB.ExecTable[updateTopicLikesCountQuery.WithArgs(1).String()] = sqlr.Result{
				Rows: 1,
			}
		}

		{
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}
		}

		likesCount, err := db.FetchBlogTopicLikeCount(context.Background(), 1)

		assert.True(t, likesCount == 1)
		assert.True(t, err == nil)

		{
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{0}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}

			stubDB.ExecTable[dislikeBlogTopicQuery.WithArgs(1, 1).String()] = sqlr.Result{
				Rows: 1,
			}
		}

		err = db.DislikeBlogTopic(context.Background(), 1, 1)

		assert.True(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, likesCount == 0)

		{
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}

			stubDB.ExecTable[likeBlogTopicQuery.WithArgs(1, 1, time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)).String()] = sqlr.Result{
				Rows: 1,
			}

			stubDB.ExecTable[updateTopicLikesCountQuery.WithArgs(1).String()] = sqlr.Result{
				Rows: 1,
			}

			stubDB.QueryTable[isBlogTopicLikedQuery.WithArgs(1, 1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}
		}

		err = db.LikeBlogTopic(context.Background(), time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC), 1, 1)

		assert.True(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, likesCount == 1)

		isLiked, err := db.IsBlogTopicLiked(context.Background(), 1, 1)

		assert.True(t, err == nil)
		assert.True(t, isLiked)
	})
}
