package db

import (
	"context"
	"testing"
	"time"

	"fantlab/assert"
	"fantlab/dbtools"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
)

func Test_FetchAvailableForums(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchAvailableForumsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,19,22"},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		availableForums, err := db.FetchAvailableForums(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, availableForums == "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,19,22")
	})

	t.Run("negative", func(t *testing.T) {
		queryTable[fetchAvailableForumsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		availableForums, err := db.FetchAvailableForums(context.Background(), 1)

		assert.True(t, availableForums == "")
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchForums(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[fetchForumsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{
				1,
				"Другие окололитературные темы",
				"Обсуждение фантастики, фэнтези и всего, что с ними связано",
				1056,
				304940,
				6144,
				"Покупка и продажа книг (только фантастика)",
				1,
				"creator",
				1,
				19,
				3206204,
				"Текст последнего сообщения",
				now,
				1,
				"Фантастика и Фэнтези",
			},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("forum_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("description"),
			dbstubs.StubColumn("topic_count"),
			dbstubs.StubColumn("message_count"),
			dbstubs.StubColumn("last_topic_id"),
			dbstubs.StubColumn("last_topic_name"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("last_message_id"),
			dbstubs.StubColumn("last_message_text"),
			dbstubs.StubColumn("last_message_date"),
			dbstubs.StubColumn("forum_block_id"),
			dbstubs.StubColumn("forum_block_name"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		forums, err := db.FetchForums(context.Background(), []uint64{1})

		assert.True(t, err == nil)
		assert.DeepEqual(t, forums, []Forum{
			{
				ForumID:         1,
				Name:            "Другие окололитературные темы",
				Description:     "Обсуждение фантастики, фэнтези и всего, что с ними связано",
				TopicCount:      1056,
				MessageCount:    304940,
				LastTopicID:     6144,
				LastTopicName:   "Покупка и продажа книг (только фантастика)",
				UserID:          1,
				Login:           "creator",
				Sex:             1,
				PhotoNumber:     19,
				LastMessageID:   3206204,
				LastMessageText: "Текст последнего сообщения",
				LastMessageDate: now,
				ForumBlockID:    1,
				ForumBlockName:  "Фантастика и Фэнтези",
			},
		})
	})

	t.Run("negative", func(t *testing.T) {
		queryTable[fetchForumsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		forums, err := db.FetchForums(context.Background(), []uint64{1})

		assert.True(t, len(forums) == 0)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchModerators(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchModeratorsQuery.String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{1, "creator", 1, 19, 2},
			{1, "creator", 1, 19, 4},
			{1, "creator", 1, 19, 11},
			{541, "vad", 1, 4, 11},
			{654, "Pirx", 1, 7, 2},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("forum_id"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		moderators, err := db.FetchModerators(context.Background())

		assert.True(t, err == nil)
		assert.DeepEqual(t, moderators, map[uint64][]ForumModerator{
			2: {
				{
					UserID:      1,
					Login:       "creator",
					Sex:         1,
					PhotoNumber: 19,
					ForumID:     2,
				},
				{
					UserID:      654,
					Login:       "Pirx",
					Sex:         1,
					PhotoNumber: 7,
					ForumID:     2,
				},
			},
			4: {
				{
					UserID:      1,
					Login:       "creator",
					Sex:         1,
					PhotoNumber: 19,
					ForumID:     4,
				},
			},
			11: {
				{
					UserID:      1,
					Login:       "creator",
					Sex:         1,
					PhotoNumber: 19,
					ForumID:     11,
				},
				{
					UserID:      541,
					Login:       "vad",
					Sex:         1,
					PhotoNumber: 4,
					ForumID:     11,
				},
			},
		})
	})

	t.Run("negative", func(t *testing.T) {
		queryTable[fetchModeratorsQuery.String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		moderators, err := db.FetchModerators(context.Background())

		assert.True(t, len(moderators) == 0)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchForumTopics(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[forumExistsQuery.WithArgs(2, []int{1, 2}).Rebind().String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	now := time.Now()

	queryTable[fetchTopicsQuery.WithArgs(2, 20, 0).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{327, "Раздел \"Авторы, книги\"", now, 519, 1, "creator", 1, 19, 1, 1, 0, 5, 10097, 203, 12, "Во-во, золотые слова:beer:", now},
			{47, "Хочу застолбить за собой автора!", now, 1969, 76, "Andrew", 1, 0, 1, 1, 0, 35, 3294, 239, 1, "Мартин Скотт, за мной:biggrin:", now},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("topic_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("date_of_add"),
			dbstubs.StubColumn("views"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("topic_type_id"),
			dbstubs.StubColumn("is_closed"),
			dbstubs.StubColumn("is_pinned"),
			dbstubs.StubColumn("message_count"),
			dbstubs.StubColumn("last_message_id"),
			dbstubs.StubColumn("last_user_id"),
			dbstubs.StubColumn("last_photo_number"),
			dbstubs.StubColumn("last_message_text"),
			dbstubs.StubColumn("last_message_date"),
		},
	}

	queryTable[topicsCountQuery.WithArgs(2).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{298}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchForumTopics(context.Background(), []uint64{1, 2}, 2, 20, 0)

		assert.True(t, err == nil)
		assert.DeepEqual(t, topics, &ForumTopicsDBResponse{
			Topics: []ForumTopic{
				{
					TopicID:         327,
					Name:            "Раздел \"Авторы, книги\"",
					DateOfAdd:       now,
					Views:           519,
					UserID:          1,
					Login:           "creator",
					Sex:             1,
					PhotoNumber:     19,
					TopicTypeID:     1,
					IsClosed:        1,
					IsPinned:        0,
					MessageCount:    5,
					LastMessageID:   10097,
					LastUserID:      203,
					LastPhotoNumber: 12,
					LastMessageText: "Во-во, золотые слова:beer:",
					LastMessageDate: now,
				},
				{
					TopicID:         47,
					Name:            "Хочу застолбить за собой автора!",
					DateOfAdd:       now,
					Views:           1969,
					UserID:          76,
					Login:           "Andrew",
					Sex:             1,
					PhotoNumber:     0,
					TopicTypeID:     1,
					IsClosed:        1,
					IsPinned:        0,
					MessageCount:    35,
					LastMessageID:   3294,
					LastUserID:      239,
					LastPhotoNumber: 1,
					LastMessageText: "Мартин Скотт, за мной:biggrin:",
					LastMessageDate: now,
				},
			},
			TotalTopicsCount: 298,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchForumTopics(context.Background(), []uint64{1, 2}, 3, 20, 0)

		assert.True(t, topics == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[topicsCountQuery.WithArgs(2).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchForumTopics(context.Background(), []uint64{1, 2}, 2, 20, 0)

		assert.True(t, topics == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		queryTable[fetchTopicsQuery.WithArgs(2, 20, 0).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchForumTopics(context.Background(), []uint64{1, 2}, 2, 20, 0)

		assert.True(t, topics == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchTopicMessages(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[shortTopicQuery.WithArgs(25, []int{1, 2}).Rebind().String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{25, "Я вас слушаю!", 2, "Техподдержка и развитие сайта"}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("topic_id"),
			dbstubs.StubColumn("topic_name"),
			dbstubs.StubColumn("forum_id"),
			dbstubs.StubColumn("forum_name"),
		},
	}

	queryTable[topicMessagesCountQuery.WithArgs(25).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{190}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	now := time.Now()

	queryTable[fetchTopicMessagesQuery.Format("ASC").WithArgs(25, 1, 20).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{216, now, 1, "creator", 1, 19, 3, "", "Ну, что скажете? По поводу сайта, естественно", 0, 0, 0},
			{224, now, 41, "Rol0c", 1, 0, 2, "", "Приветик ;-)", 0, 0, 0},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("message_id"),
			dbstubs.StubColumn("date_of_add"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("user_class"),
			dbstubs.StubColumn("sign"),
			dbstubs.StubColumn("message_text"),
			dbstubs.StubColumn("is_censored"),
			dbstubs.StubColumn("vote_plus"),
			dbstubs.StubColumn("vote_minus"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchTopicMessages(context.Background(), []uint64{1, 2}, 25, 20, 0, "ASC")

		assert.True(t, err == nil)
		assert.DeepEqual(t, topics, &ForumTopicMessagesDBResponse{
			Topic: ShortForumTopic{
				TopicID:   25,
				TopicName: "Я вас слушаю!",
				ForumID:   2,
				ForumName: "Техподдержка и развитие сайта",
			},
			Messages: []ForumMessage{
				{
					MessageID:   216,
					DateOfAdd:   now,
					UserID:      1,
					Login:       "creator",
					Sex:         1,
					PhotoNumber: 19,
					UserClass:   3,
					Sign:        "",
					MessageText: "Ну, что скажете? По поводу сайта, естественно",
					IsCensored:  0,
					VotePlus:    0,
					VoteMinus:   0,
				},
				{
					MessageID:   224,
					DateOfAdd:   now,
					UserID:      41,
					Login:       "Rol0c",
					Sex:         1,
					PhotoNumber: 0,
					UserClass:   2,
					Sign:        "",
					MessageText: "Приветик ;-)",
					IsCensored:  0,
					VotePlus:    0,
					VoteMinus:   0,
				},
			},
			TotalMessagesCount: 190,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchTopicMessages(context.Background(), []uint64{1, 2}, 1, 20, 0, "ASC")

		assert.True(t, topics == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[fetchTopicMessagesQuery.Format("ASC").WithArgs(25, 1, 20).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchTopicMessages(context.Background(), []uint64{1, 2}, 25, 20, 0, "ASC")

		assert.True(t, topics == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		queryTable[topicMessagesCountQuery.WithArgs(25).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		topics, err := db.FetchTopicMessages(context.Background(), []uint64{1, 2}, 25, 20, 0, "ASC")

		assert.True(t, topics == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})
}
