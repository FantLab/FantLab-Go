package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
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
			stubDB.ExecTable[sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(1).String()] = sqlr.Result{
				Rows: 1,
			}
		}

		{
			stubDB.QueryTable[sqlr.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(1).String()] = &dbstubs.StubRows{
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
			stubDB.QueryTable[sqlr.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{0}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}

			stubDB.ExecTable[sqlr.NewQuery(queries.DislikeBlogTopic).WithArgs(1, 1).String()] = sqlr.Result{
				Rows: 1,
			}
		}

		err = db.DislikeBlogTopic(context.Background(), 1, 1)

		assert.True(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, likesCount == 0)

		{
			stubDB.QueryTable[sqlr.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}

			stubDB.ExecTable[sqlr.NewQuery(queries.LikeBlogTopic).WithArgs(1, 1, time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)).String()] = sqlr.Result{
				Rows: 1,
			}

			stubDB.ExecTable[sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(1).String()] = sqlr.Result{
				Rows: 1,
			}

			stubDB.QueryTable[sqlr.NewQuery(queries.IsBlogTopicLiked).WithArgs(1, 1).String()] = &dbstubs.StubRows{
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
