package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
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
				RowsAffected: 1,
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
				RowsAffected: 1,
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

			stubDB.ExecTable[sqlr.NewQuery(queries.LikeBlogTopic).WithArgs(1, 1).String()] = sqlr.Result{
				RowsAffected: 1,
			}

			stubDB.ExecTable[sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(1).String()] = sqlr.Result{
				RowsAffected: 1,
			}
		}

		err = db.LikeBlogTopic(context.Background(), 1, 1)

		assert.True(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, likesCount == 1)
	})
}
