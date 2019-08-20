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
			stubDB.ExecTable["UPDATE b_topics b SET b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id) WHERE b.topic_id = 1"] = sqlr.Result{
				Rows: 1,
			}
		}

		{
			stubDB.QueryTable["SELECT likes_count FROM b_topics WHERE topic_id = 1"] = &stubs.StubRows{
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
			stubDB.QueryTable["SELECT likes_count FROM b_topics WHERE topic_id = 1"] = &stubs.StubRows{
				Values: [][]interface{}{{0}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
				},
			}

			stubDB.ExecTable["DELETE FROM b_topic_likes WHERE topic_id = 1 AND user_id = 1"] = sqlr.Result{
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
			stubDB.QueryTable["SELECT likes_count FROM b_topics WHERE topic_id = 1"] = &stubs.StubRows{
				Values: [][]interface{}{{1}},
				Columns: []scanr.Column{
					stubs.StubColumn(""),
				},
			}

			stubDB.ExecTable["INSERT INTO b_topic_likes (topic_id, user_id, date_of_add) VALUES (1, 1, '2019-08-19 17:40:03')"] = sqlr.Result{
				Rows: 1,
			}

			stubDB.QueryTable["SELECT 1 FROM b_topic_likes WHERE topic_id = 1 AND user_id = 1"] = &stubs.StubRows{
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
