package db

import (
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"fantlab/dbtools/stubs"
	"fantlab/tt"
	"testing"
	"time"
)

func Test_LikeDiskie(t *testing.T) {
	stubDB := &stubs.StubDB{
		ExecTable:  make(stubs.StubExecTable),
		QueryTable: make(stubs.StubQueryTable),
	}

	db := &DB{R: stubDB}

	t.Run("like_dislike", func(t *testing.T) {
		{
			stubDB.ExecTable[updateTopicLikesCountQuery.WithArgs(1).String()] = sqlr.Result{
				Rows: 1,
			}
		}

		{
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &stubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
				},
			}
		}

		likesCount, err := db.FetchBlogTopicLikeCount(1)

		tt.Assert(t, likesCount == 1)
		tt.Assert(t, err == nil)

		{
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &stubs.StubRows{
				Values: [][]interface{}{{0}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
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
			stubDB.QueryTable[fetchBlogTopicLikeCountQuery.WithArgs(1).String()] = &stubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
				},
			}

			stubDB.ExecTable[likeBlogTopicQuery.WithArgs(1, 1, time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)).String()] = sqlr.Result{
				Rows: 1,
			}

			stubDB.QueryTable[isBlogTopicLikedQuery.WithArgs(1, 1).String()] = &stubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
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
