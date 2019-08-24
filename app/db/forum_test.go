package db

import (
	"testing"
	"time"

	"fantlab/dbtools/scanr"
	"fantlab/dbtools/stubs"
	"fantlab/tt"
)

func Test_FetchForums(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	now := time.Now()

	queryTable[fetchForumsQuery.WithArgs(1).String()] = &stubs.StubRows{
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
			stubs.StubColumn("forum_id"),
			stubs.StubColumn("name"),
			stubs.StubColumn("description"),
			stubs.StubColumn("topic_count"),
			stubs.StubColumn("message_count"),
			stubs.StubColumn("last_topic_id"),
			stubs.StubColumn("last_topic_name"),
			stubs.StubColumn("user_id"),
			stubs.StubColumn("login"),
			stubs.StubColumn("sex"),
			stubs.StubColumn("photo_number"),
			stubs.StubColumn("last_message_id"),
			stubs.StubColumn("last_message_text"),
			stubs.StubColumn("last_message_date"),
			stubs.StubColumn("forum_block_id"),
			stubs.StubColumn("forum_block_name"),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		forums, err := db.FetchForums([]uint16{1})

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, forums, []Forum{
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
}

func Test_FetchModerators(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	queryTable[fetchModeratorsQuery.String()] = &stubs.StubRows{
		Values: [][]interface{}{
			{1, "creator", 1, 19, 2},
			{1, "creator", 1, 19, 4},
			{1, "creator", 1, 19, 11},
			{541, "vad", 1, 4, 11},
			{654, "Pirx", 1, 7, 2},
		},
		Columns: []scanr.Column{
			stubs.StubColumn("user_id"),
			stubs.StubColumn("login"),
			stubs.StubColumn("sex"),
			stubs.StubColumn("photo_number"),
			stubs.StubColumn("forum_id"),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		moderators, err := db.FetchModerators()

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, moderators, map[uint32][]ForumModerator{
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
}
