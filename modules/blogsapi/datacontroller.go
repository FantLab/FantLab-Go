package blogsapi

import (
	"fantlab/protobuf/generated/fantlab/apimodels"
	"fantlab/utils"
)

func getCommunities(dbCommunities []dbCommunity) *apimodels.Blog_CommunitiesResponse {
	var mainCommunities []*apimodels.Blog_Community
	var additionalCommunities []*apimodels.Blog_Community

	for _, dbCommunity := range dbCommunities {
		community := &apimodels.Blog_Community{
			Id:                   dbCommunity.BlogId,
			Title:                dbCommunity.Name,
			CommunityDescription: dbCommunity.Description,
			Stats: &apimodels.Blog_Community_Stats{
				ArticlesCount:    dbCommunity.TopicsCount,
				SubscribersCount: dbCommunity.SubscriberCount,
			},
			LastArticle: &apimodels.Blog_LastArticle{
				Id:    dbCommunity.LastTopicId,
				Title: dbCommunity.LastTopicHead,
				User: &apimodels.Blog_UserLink{
					Id:    dbCommunity.LastUserId,
					Login: dbCommunity.LastUserName,
				},
				Date: utils.ProtoTS(dbCommunity.LastTopicDate),
			},
		}

		if dbCommunity.IsPublic {
			mainCommunities = append(mainCommunities, community)
		} else {
			additionalCommunities = append(additionalCommunities, community)
		}
	}

	return &apimodels.Blog_CommunitiesResponse{
		Main:       mainCommunities,
		Additional: additionalCommunities,
	}
}

func getBlogs(dbBlogs []dbBlog) *apimodels.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*apimodels.Blog_Blog{}

	for _, dbBlog := range dbBlogs {
		blog := &apimodels.Blog_Blog{
			Id: dbBlog.BlogId,
			User: &apimodels.Blog_UserLink{
				Id:    dbBlog.UserId,
				Login: dbBlog.Login,
				Name:  dbBlog.Fio,
			},
			IsClosed: dbBlog.IsClose,
			Stats: &apimodels.Blog_Blog_Stats{
				ArticlesCount:    dbBlog.TopicsCount,
				SubscribersCount: dbBlog.SubscriberCount,
			},
			LastArticle: &apimodels.Blog_LastArticle{
				Id:    dbBlog.LastTopicId,
				Title: dbBlog.LastTopicHead,
				Date:  utils.ProtoTS(dbBlog.LastTopicDate),
			},
		}

		blogs = append(blogs, blog)
	}

	return &apimodels.Blog_BlogsResponse{
		Blogs: blogs,
	}
}

func getBlogArticles(dbBlogTopics []dbBlogTopic, urlFormatter utils.UrlFormatter) *apimodels.Blog_ArticlesResponse {
	//noinspection GoPreferNilSlice
	var articles = []*apimodels.Blog_Article{}

	for _, dbBlogTopic := range dbBlogTopics {
		var gender apimodels.Gender
		if dbBlogTopic.Sex == 0 {
			gender = apimodels.Gender_FEMALE
		} else {
			gender = apimodels.Gender_MALE
		}

		avatar := urlFormatter.GetAvatarUrl(dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

		article := &apimodels.Blog_Article{
			Id:    dbBlogTopic.TopicId,
			Title: dbBlogTopic.HeadTopic,
			Creation: &apimodels.Blog_Creation{
				User: &apimodels.Blog_UserLink{
					Id:     dbBlogTopic.UserId,
					Login:  dbBlogTopic.Login,
					Gender: gender,
					Avatar: avatar,
				},
				Date: utils.ProtoTS(dbBlogTopic.DateOfAdd),
			},
			Text: dbBlogTopic.MessageText,
			Tags: dbBlogTopic.Tags,
			Stats: &apimodels.Blog_Article_Stats{
				LikeCount:    dbBlogTopic.LikesCount,
				ViewCount:    dbBlogTopic.Views,
				CommentCount: dbBlogTopic.CommentsCount,
			},
		}

		articles = append(articles, article)
	}

	return &apimodels.Blog_ArticlesResponse{
		Articles: articles,
	}
}
