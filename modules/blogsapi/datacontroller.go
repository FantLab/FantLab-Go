package blogsapi

func getCommunities(dbCommunities []dbCommunity) communitiesWrapper {
	var mainCommunities []community
	var additionalCommunities []community

	for _, dbCommunity := range dbCommunities {
		community := community{
			Id:          dbCommunity.BlogId,
			Title:       dbCommunity.Name,
			Description: dbCommunity.Description,
			Stats: stats{
				ArticleCount:    dbCommunity.TopicsCount,
				SubscriberCount: dbCommunity.SubscriberCount,
			},
			LastArticle: lastArticle{
				Id:    dbCommunity.LastTopicId,
				Title: dbCommunity.LastTopicHead,
				User: &userLink{
					Id:    dbCommunity.LastUserId,
					Login: dbCommunity.LastUserName,
				},
				Date: dbCommunity.LastTopicDate.Unix(),
			},
		}
		if dbCommunity.IsPublic {
			mainCommunities = append(mainCommunities, community)
		} else {
			additionalCommunities = append(additionalCommunities, community)
		}
	}

	return communitiesWrapper{
		Main:       mainCommunities,
		Additional: additionalCommunities,
	}
}

func getBlogs(dbBlogs []dbBlog) blogsWrapper {
	//noinspection GoPreferNilSlice
	var blogs = []blog{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbBlog := range dbBlogs {
		blog := blog{
			Id: dbBlog.BlogId,
			Owner: userLink{
				Id:    dbBlog.UserId,
				Login: dbBlog.Login,
				Name:  dbBlog.Fio,
			},
			IsClosed: dbBlog.IsClose,
			Stats: stats{
				ArticleCount:    dbBlog.TopicsCount,
				SubscriberCount: dbBlog.SubscriberCount,
			},
			LastArticle: lastArticle{
				Id:    dbBlog.LastTopicId,
				Title: dbBlog.LastTopicHead,
				Date:  dbBlog.LastTopicDate.Unix(),
			},
		}
		blogs = append(blogs, blog)
	}

	return blogsWrapper{blogs}
}

func getBlogArticles(dbBlogTopics []dbBlogTopic) blogArticlesWrapper {
	//noinspection GoPreferNilSlice
	var articles = []article{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbBlogTopic := range dbBlogTopics {
		article := article{
			Id:    dbBlogTopic.TopicId,
			Title: dbBlogTopic.HeadTopic,
			Creation: creation{
				User: userLink{
					Id:    dbBlogTopic.UserId,
					Login: dbBlogTopic.Login,
				},
				Date: dbBlogTopic.DateOfAdd.Unix(),
			},
			Text: dbBlogTopic.MessageText,
			Tags: dbBlogTopic.Tags,
			Stats: articleStats{
				LikeCount:    dbBlogTopic.LikesCount,
				ViewCount:    dbBlogTopic.Views,
				CommentCount: dbBlogTopic.CommentsCount,
			},
		}
		articles = append(articles, article)
	}

	return blogArticlesWrapper{articles}
}
