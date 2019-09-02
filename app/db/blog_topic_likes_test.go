package db

import (
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"fantlab/tt"
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

		likesCount, err := db.FetchBlogTopicLikeCount(1)

		tt.Assert(t, likesCount == 1)
		tt.Assert(t, err == nil)

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

		ok, err := db.DislikeBlogTopic(1, 1)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(1)

		tt.Assert(t, err == nil)
		tt.Assert(t, likesCount == 0)

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

			stubDB.QueryTable[isBlogTopicLikedQuery.WithArgs(1, 1).String()] = &dbstubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					dbstubs.StubColumn(""),
				},
			}
		}

		ok, err = db.LikeBlogTopic(time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC), 1, 1)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)

		likesCount, err = db.FetchBlogTopicLikeCount(1)

		tt.Assert(t, err == nil)
		tt.Assert(t, likesCount == 1)

		isLiked, err := db.IsBlogTopicLiked(1, 1)

		tt.Assert(t, err == nil)
		tt.Assert(t, isLiked)
	})
}
