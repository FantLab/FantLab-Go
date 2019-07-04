package blogsapi

import (
	"fantlab/pb"
	"fantlab/utils"
)

func getCommunities(dbCommunities []dbCommunity, urlFormatter utils.UrlFormatter) *pb.Blog_CommunitiesResponse {
	var mainCommunities []*pb.Blog_Community
	var additionalCommunities []*pb.Blog_Community

	for _, dbCommunity := range dbCommunities {
		gender := utils.GetGender(dbCommunity.LastSex)
		avatar := urlFormatter.GetAvatarUrl(dbCommunity.LastUserId, dbCommunity.LastPhotoNumber)

		community := &pb.Blog_Community{
			Id:                   dbCommunity.BlogId,
			Title:                dbCommunity.Name,
			CommunityDescription: dbCommunity.Description,
			Stats: &pb.Blog_Community_Stats{
				ArticleCount:    dbCommunity.TopicsCount,
				SubscriberCount: dbCommunity.SubscriberCount,
			},
			LastArticle: &pb.Blog_LastArticle{
				Id:    dbCommunity.LastTopicId,
				Title: dbCommunity.LastTopicHead,
				User: &pb.Common_UserLink{
					Id:     dbCommunity.LastUserId,
					Login:  dbCommunity.LastLogin,
					Gender: gender,
					Avatar: avatar,
				},
				Text: dbCommunity.LastTopicText,
				Date: utils.ProtoTS(dbCommunity.LastTopicDate),
			},
		}

		if dbCommunity.IsPublic {
			mainCommunities = append(mainCommunities, community)
		} else {
			additionalCommunities = append(additionalCommunities, community)
		}
	}

	return &pb.Blog_CommunitiesResponse{
		Main:       mainCommunities,
		Additional: additionalCommunities,
	}
}

func getBlogs(dbBlogs []dbBlog, urlFormatter utils.UrlFormatter) *pb.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*pb.Blog_Blog{}

	for _, dbBlog := range dbBlogs {
		gender := utils.GetGender(dbBlog.Sex)
		avatar := urlFormatter.GetAvatarUrl(dbBlog.UserId, dbBlog.PhotoNumber)

		blog := &pb.Blog_Blog{
			Id: dbBlog.BlogId,
			User: &pb.Common_UserLink{
				Id:     dbBlog.UserId,
				Login:  dbBlog.Login,
				Name:   dbBlog.Fio,
				Gender: gender,
				Avatar: avatar,
			},
			IsClosed: dbBlog.IsClose,
			Stats: &pb.Blog_Blog_Stats{
				ArticleCount:    dbBlog.TopicsCount,
				SubscriberCount: dbBlog.SubscriberCount,
			},
			LastArticle: &pb.Blog_LastArticle{
				Id:    dbBlog.LastTopicId,
				Title: dbBlog.LastTopicHead,
				Text:  dbBlog.LastTopicText,
				Date:  utils.ProtoTS(dbBlog.LastTopicDate),
			},
		}

		blogs = append(blogs, blog)
	}

	return &pb.Blog_BlogsResponse{
		Blogs: blogs,
	}
}

func getBlogArticles(dbBlogTopics []dbBlogTopic, urlFormatter utils.UrlFormatter) *pb.Blog_ArticlesResponse {
	//noinspection GoPreferNilSlice
	var articles = []*pb.Blog_Article{}

	for _, dbBlogTopic := range dbBlogTopics {
		gender := utils.GetGender(dbBlogTopic.Sex)
		avatar := urlFormatter.GetAvatarUrl(dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

		article := &pb.Blog_Article{
			Id:    dbBlogTopic.TopicId,
			Title: dbBlogTopic.HeadTopic,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbBlogTopic.UserId,
					Login:  dbBlogTopic.Login,
					Gender: gender,
					Avatar: avatar,
				},
				Date: utils.ProtoTS(dbBlogTopic.DateOfAdd),
			},
			Text: dbBlogTopic.MessageText,
			Tags: dbBlogTopic.Tags,
			Stats: &pb.Blog_Article_Stats{
				LikeCount:    dbBlogTopic.LikesCount,
				ViewCount:    dbBlogTopic.Views,
				CommentCount: dbBlogTopic.CommentsCount,
			},
		}

		articles = append(articles, article)
	}

	return &pb.Blog_ArticlesResponse{
		Articles: articles,
	}
}
