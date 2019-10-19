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

func Test_FetchCommunities(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[communitiesQuery.String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{271, "КИНОрецензии", "рецензии на фильмы", 635, 1, now, "Первому игроку приготовиться", 55779, 362, 7039, "Phelan", 1, 2},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("blog_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("description"),
			dbstubs.StubColumn("topics_count"),
			dbstubs.StubColumn("is_public"),
			dbstubs.StubColumn("last_topic_date"),
			dbstubs.StubColumn("last_topic_head"),
			dbstubs.StubColumn("last_topic_id"),
			dbstubs.StubColumn("subscriber_count"),
			dbstubs.StubColumn("last_user_id"),
			dbstubs.StubColumn("last_login"),
			dbstubs.StubColumn("last_sex"),
			dbstubs.StubColumn("last_photo_number"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communities, err := db.FetchCommunities(context.Background())

		assert.True(t, err == nil)
		assert.DeepEqual(t, communities, []Community{
			{
				BlogId:          271,
				Name:            "КИНОрецензии",
				Description:     "рецензии на фильмы",
				TopicsCount:     635,
				IsPublic:        1,
				LastTopicDate:   now,
				LastTopicHead:   "Первому игроку приготовиться",
				LastTopicId:     55779,
				SubscriberCount: 362,
				LastUserId:      7039,
				LastLogin:       "Phelan",
				LastSex:         1,
				LastPhotoNumber: 2,
			},
		})
	})

	t.Run("negative", func(t *testing.T) {
		queryTable[communitiesQuery.String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communities, err := db.FetchCommunities(context.Background())

		assert.True(t, len(communities) == 0)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchCommunity(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[communityQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{"КИНОрецензии", "рецензии на фильмы", 635, 1, now, "Первому игроку приготовиться", 55779, 362, 7039, "Phelan", 1, 2},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("description"),
			dbstubs.StubColumn("topics_count"),
			dbstubs.StubColumn("is_public"),
			dbstubs.StubColumn("last_topic_date"),
			dbstubs.StubColumn("last_topic_head"),
			dbstubs.StubColumn("last_topic_id"),
			dbstubs.StubColumn("subscriber_count"),
			dbstubs.StubColumn("last_user_id"),
			dbstubs.StubColumn("last_login"),
			dbstubs.StubColumn("last_sex"),
			dbstubs.StubColumn("last_photo_number"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		community, err := db.FetchCommunity(context.Background(), 271)

		assert.True(t, err == nil)
		assert.DeepEqual(t, community, &Community{
			Name:            "КИНОрецензии",
			Description:     "рецензии на фильмы",
			TopicsCount:     635,
			IsPublic:        1,
			LastTopicDate:   now,
			LastTopicHead:   "Первому игроку приготовиться",
			LastTopicId:     55779,
			SubscriberCount: 362,
			LastUserId:      7039,
			LastLogin:       "Phelan",
			LastSex:         1,
			LastPhotoNumber: 2,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		community, err := db.FetchCommunity(context.Background(), 1)

		assert.True(t, community == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})
}

func Test_FetchCommunityTopics(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[shortCommunityQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{271, "КИНОрецензии", "Рубрика для отобранных и качественных рецензий на кинофильмы."},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("blog_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("rules"),
		},
	}

	queryTable[communityModeratorsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{651, "Barros", 1, 1},
			{8775, "fox_mulder", 1, 1237},
			{3995, "sham", 1, 2},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
		},
	}

	queryTable[communityAuthorsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{5396, "glupec", 1, 28},
			{7246, "armitura", 1, 54},
			{15193, "alexsei111", 1, 4},
			{16934, "rusty_cat", 1, 9},
			{14228, "Mierin", 0, 13},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
		},
	}

	now := time.Now()

	queryTable[communityTopicsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{12387, "Плывет корабль...", now, 5396, "glupec", 1, 28, "Текст статьи 12387", "Клайв С. Льюис", 8},
			{11260, "Не ТРОНь !", now, 7999, "febeerovez", 1, 52, "Текст статьи 11260", "Трон: Наследие", 46},
			{10272, "Книга Илая", now, 20874, "Fadvan", 1, 19, "Текст статьи 10272", "Книга Илая", 37},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("topic_id"),
			dbstubs.StubColumn("head_topic"),
			dbstubs.StubColumn("date_of_add"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("message_text"),
			dbstubs.StubColumn("tags"),
			dbstubs.StubColumn("comments_count"),
		},
	}

	queryTable[communityTopicCountQuery.WithArgs(271).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{9}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 271, 3, 0)

		assert.True(t, err == nil)
		assert.DeepEqual(t, communityTopicsDbResponse, &CommunityTopicsDBResponse{
			Community: Community{
				BlogId: 271,
				Name:   "КИНОрецензии",
				Rules:  "Рубрика для отобранных и качественных рецензий на кинофильмы.",
			},
			Moderators: []CommunityModerator{
				{
					UserID:      651,
					Login:       "Barros",
					Sex:         1,
					PhotoNumber: 1,
				},
				{
					UserID:      8775,
					Login:       "fox_mulder",
					Sex:         1,
					PhotoNumber: 1237,
				},
				{
					UserID:      3995,
					Login:       "sham",
					Sex:         1,
					PhotoNumber: 2,
				},
			},
			Authors: []CommunityAuthor{
				{
					UserID:      5396,
					Login:       "glupec",
					Sex:         1,
					PhotoNumber: 28,
				},
				{
					UserID:      7246,
					Login:       "armitura",
					Sex:         1,
					PhotoNumber: 54,
				},
				{
					UserID:      15193,
					Login:       "alexsei111",
					Sex:         1,
					PhotoNumber: 4,
				},
				{
					UserID:      16934,
					Login:       "rusty_cat",
					Sex:         1,
					PhotoNumber: 9,
				},
				{
					UserID:      14228,
					Login:       "Mierin",
					Sex:         0,
					PhotoNumber: 13,
				},
			},
			Topics: []BlogTopic{
				{
					TopicId:       12387,
					HeadTopic:     "Плывет корабль...",
					DateOfAdd:     now,
					UserId:        5396,
					Login:         "glupec",
					Sex:           1,
					PhotoNumber:   28,
					MessageText:   "Текст статьи 12387",
					Tags:          "Клайв С. Льюис",
					CommentsCount: 8,
				},
				{
					TopicId:       11260,
					HeadTopic:     "Не ТРОНь !",
					DateOfAdd:     now,
					UserId:        7999,
					Login:         "febeerovez",
					Sex:           1,
					PhotoNumber:   52,
					MessageText:   "Текст статьи 11260",
					Tags:          "Трон: Наследие",
					CommentsCount: 46,
				},
				{
					TopicId:       10272,
					HeadTopic:     "Книга Илая",
					DateOfAdd:     now,
					UserId:        20874,
					Login:         "Fadvan",
					Sex:           1,
					PhotoNumber:   19,
					MessageText:   "Текст статьи 10272",
					Tags:          "Книга Илая",
					CommentsCount: 37,
				},
			},
			TotalTopicsCount: 9,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 1, 3, 0)

		assert.True(t, communityTopicsDbResponse == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_1", func(t *testing.T) {
		queryTable[communityTopicCountQuery.WithArgs(271).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 271, 3, 0)

		assert.True(t, communityTopicsDbResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[communityTopicsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 271, 3, 0)

		assert.True(t, communityTopicsDbResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		queryTable[communityAuthorsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 271, 3, 0)

		assert.True(t, communityTopicsDbResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_4", func(t *testing.T) {
		queryTable[communityModeratorsQuery.WithArgs(271).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		communityTopicsDbResponse, err := db.FetchCommunityTopics(context.Background(), 271, 3, 0)

		assert.True(t, communityTopicsDbResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchBlogs(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[blogsQuery.Format("subscriber_count").WithArgs(5, 0).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{229, 10072, "antilia", "Гарбузова Ольга", 0, 2, 436, 975, 0, now, "Книжные новинки за неделю", 55717},
			{611, 17299, "k2007", "", 1, 14, 341, 597, 0, now, "Нефантастика, детская фантастическая литература", 55745},
			{112, 3068, "Vladimir Puziy", "", 1, 0, 723, 583, 0, now, "Презентация нового издания", 55690},
			{33, 143, "Dark Andrew", "Андрей Зильберштейн", 1, 13, 154, 541, 0, now, "Раздумчивое о розах и червях", 48217},
			{408, 13240, "С.Соболев", "Соболев Сергей Васильевич", 1, 1, 177, 454, 0, now, "На всё лето - в поля", 55405},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("blog_id"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("fio"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("topics_count"),
			dbstubs.StubColumn("subscriber_count"),
			dbstubs.StubColumn("is_close"),
			dbstubs.StubColumn("last_topic_date"),
			dbstubs.StubColumn("last_topic_head"),
			dbstubs.StubColumn("last_topic_id"),
		},
	}

	queryTable[blogCountQuery.String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{645}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogsDbResponse, err := db.FetchBlogs(context.Background(), 5, 0, "subscriber")

		assert.True(t, err == nil)
		assert.DeepEqual(t, blogsDbResponse, &BlogsDBResponse{
			Blogs: []Blog{
				{
					BlogId:          229,
					UserId:          10072,
					Login:           "antilia",
					Fio:             "Гарбузова Ольга",
					Sex:             0,
					PhotoNumber:     2,
					TopicsCount:     436,
					SubscriberCount: 975,
					IsClose:         0,
					LastTopicDate:   now,
					LastTopicHead:   "Книжные новинки за неделю",
					LastTopicId:     55717,
				},
				{
					BlogId:          611,
					UserId:          17299,
					Login:           "k2007",
					Fio:             "",
					Sex:             1,
					PhotoNumber:     14,
					TopicsCount:     341,
					SubscriberCount: 597,
					IsClose:         0,
					LastTopicDate:   now,
					LastTopicHead:   "Нефантастика, детская фантастическая литература",
					LastTopicId:     55745,
				},
				{
					BlogId:          112,
					UserId:          3068,
					Login:           "Vladimir Puziy",
					Fio:             "",
					Sex:             1,
					PhotoNumber:     0,
					TopicsCount:     723,
					SubscriberCount: 583,
					IsClose:         0,
					LastTopicDate:   now,
					LastTopicHead:   "Презентация нового издания",
					LastTopicId:     55690,
				},
				{
					BlogId:          33,
					UserId:          143,
					Login:           "Dark Andrew",
					Fio:             "Андрей Зильберштейн",
					Sex:             1,
					PhotoNumber:     13,
					TopicsCount:     154,
					SubscriberCount: 541,
					IsClose:         0,
					LastTopicDate:   now,
					LastTopicHead:   "Раздумчивое о розах и червях",
					LastTopicId:     48217,
				},
				{
					BlogId:          408,
					UserId:          13240,
					Login:           "С.Соболев",
					Fio:             "Соболев Сергей Васильевич",
					Sex:             1,
					PhotoNumber:     1,
					TopicsCount:     177,
					SubscriberCount: 454,
					IsClose:         0,
					LastTopicDate:   now,
					LastTopicHead:   "На всё лето - в поля",
					LastTopicId:     55405,
				},
			},
			TotalCount: 645,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogsDbResponse, err := db.FetchBlogs(context.Background(), 5, 0, "update")

		assert.True(t, blogsDbResponse == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[blogCountQuery.String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogsDbResponse, err := db.FetchBlogs(context.Background(), 5, 0, "subscriber")

		assert.True(t, blogsDbResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchBlog(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[blogQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{1, 1, "creator", "Львов Алексей", 1, 19, 22, 112, 0, now, "Майкл Ши. Первые экземпляры", 55320},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("blog_id"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("fio"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("topics_count"),
			dbstubs.StubColumn("subscriber_count"),
			dbstubs.StubColumn("is_close"),
			dbstubs.StubColumn("last_topic_date"),
			dbstubs.StubColumn("last_topic_head"),
			dbstubs.StubColumn("last_topic_id"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blog, err := db.FetchBlog(context.Background(), 1)

		assert.True(t, err == nil)
		assert.DeepEqual(t, blog, &Blog{
			BlogId:          1,
			UserId:          1,
			Login:           "creator",
			Fio:             "Львов Алексей",
			Sex:             1,
			PhotoNumber:     19,
			TopicsCount:     22,
			SubscriberCount: 112,
			IsClose:         0,
			LastTopicDate:   now,
			LastTopicHead:   "Майкл Ши. Первые экземпляры",
			LastTopicId:     55320,
		})
	})

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blog, err := db.FetchBlog(context.Background(), 2)

		assert.True(t, blog == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})
}

func Test_FetchBlogTopics(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[blogExistsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	now := time.Now()

	queryTable[blogTopicsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{52, "Авторские колонки", now, 1, "creator", 1, 19, "Новый раздел авторских колонок...", "авторские колонки", 29},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("topic_id"),
			dbstubs.StubColumn("head_topic"),
			dbstubs.StubColumn("date_of_add"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("message_text"),
			dbstubs.StubColumn("tags"),
			dbstubs.StubColumn("comments_count"),
		},
	}

	queryTable[blogTopicCountQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{5}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopicsDBResponse, err := db.FetchBlogTopics(context.Background(), 1, 1, 0)

		assert.True(t, err == nil)
		assert.DeepEqual(t, blogTopicsDBResponse, &BlogTopicsDBResponse{
			Topics: []BlogTopic{
				{
					TopicId:       52,
					HeadTopic:     "Авторские колонки",
					DateOfAdd:     now,
					UserId:        1,
					Login:         "creator",
					Sex:           1,
					PhotoNumber:   19,
					MessageText:   "Новый раздел авторских колонок...",
					Tags:          "авторские колонки",
					CommentsCount: 29,
				},
			},
			TotalTopicsCount: 5,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopicsDBResponse, err := db.FetchBlogTopics(context.Background(), 2, 1, 0)

		assert.True(t, blogTopicsDBResponse == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[blogTopicCountQuery.WithArgs(1).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopicsDBResponse, err := db.FetchBlogTopics(context.Background(), 1, 1, 0)

		assert.True(t, blogTopicsDBResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[blogTopicsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopicsDBResponse, err := db.FetchBlogTopics(context.Background(), 1, 1, 0)

		assert.True(t, blogTopicsDBResponse == nil)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchBlogTopic(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	now := time.Now()

	queryTable[topicQuery.WithArgs(52).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{52, "Авторские колонки", now, 1, "creator", 1, 19, "Новый раздел авторских колонок...", "авторские колонки", 29},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("topic_id"),
			dbstubs.StubColumn("head_topic"),
			dbstubs.StubColumn("date_of_add"),
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("photo_number"),
			dbstubs.StubColumn("message_text"),
			dbstubs.StubColumn("tags"),
			dbstubs.StubColumn("comments_count"),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopic, err := db.FetchBlogTopic(context.Background(), 52)

		assert.True(t, err == nil)
		assert.DeepEqual(t, blogTopic, &BlogTopic{
			TopicId:       52,
			HeadTopic:     "Авторские колонки",
			DateOfAdd:     now,
			UserId:        1,
			Login:         "creator",
			Sex:           1,
			PhotoNumber:   19,
			MessageText:   "Новый раздел авторских колонок...",
			Tags:          "авторские колонки",
			CommentsCount: 29,
		})
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		blogTopic, err := db.FetchBlogTopic(context.Background(), 1)

		assert.True(t, blogTopic == nil)
		assert.True(t, dbtools.IsNotFoundError(err))
	})
}
