package blogsapi

import (
	"fantlab/utils"

	"strconv"
)

func getCommunities(dbCommunities []dbCommunity, isDebug bool) communitiesWrapper {
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

		if isDebug {
			lastArticleDebugDate := utils.FormatDebugTime(dbCommunity.LastTopicDate)
			community.LastArticle.DebugDate = &lastArticleDebugDate
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

func getBlogs(dbBlogs []dbBlog, isDebug bool) blogsWrapper {
	//noinspection GoPreferNilSlice
	var blogs = []blog{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbBlog := range dbBlogs {
		blog := blog{
			Id: dbBlog.BlogId,
			User: userLink{
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

		if isDebug {
			lastArticleDebugDate := utils.FormatDebugTime(dbBlog.LastTopicDate)
			blog.LastArticle.DebugDate = &lastArticleDebugDate
		}

		blogs = append(blogs, blog)
	}

	return blogsWrapper{blogs}
}

func getBlogArticles(dbBlogTopics []dbBlogTopic, imageUrl string, isDebug bool) blogArticlesWrapper {
	//noinspection GoPreferNilSlice
	var articles = []article{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbBlogTopic := range dbBlogTopics {
		var gender string
		if dbBlogTopic.Sex == 0 {
			gender = "f"
		} else {
			gender = "m"
		}

		var avatar string
		if dbBlogTopic.PhotoNumber != 0 {
			userId := strconv.FormatUint(uint64(dbBlogTopic.UserId), 10)
			photoNumber := strconv.FormatUint(uint64(dbBlogTopic.PhotoNumber), 10)
			avatar = imageUrl + "/users/" + userId + "_" + photoNumber
		}

		article := article{
			Id:    dbBlogTopic.TopicId,
			Title: dbBlogTopic.HeadTopic,
			Creation: creation{
				User: userLink{
					Id:     dbBlogTopic.UserId,
					Login:  dbBlogTopic.Login,
					Gender: gender,
					Avatar: avatar,
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

		if isDebug {
			creationDebugDate := utils.FormatDebugTime(dbBlogTopic.DateOfAdd)
			article.Creation.DebugDate = &creationDebugDate
		}

		articles = append(articles, article)
	}

	return blogArticlesWrapper{articles}
}
