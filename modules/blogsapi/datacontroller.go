package blogsapi

import (
	"fantlab/protobuf/generated/fantlab/pb"
	"fantlab/utils"
)

func getCommunities(dbCommunities []dbCommunity) *pb.Blog_CommunitiesResponse {
	var mainCommunities []*pb.Blog_Community
	var additionalCommunities []*pb.Blog_Community

	for _, dbCommunity := range dbCommunities {
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
				User: &pb.Blog_UserLink{
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

	return &pb.Blog_CommunitiesResponse{
		Main:       mainCommunities,
		Additional: additionalCommunities,
	}
}

func getBlogs(dbBlogs []dbBlog) *pb.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*pb.Blog_Blog{}

	for _, dbBlog := range dbBlogs {
		blog := &pb.Blog_Blog{
			Id: dbBlog.BlogId,
			User: &pb.Blog_UserLink{
				Id:    dbBlog.UserId,
				Login: dbBlog.Login,
				Name:  dbBlog.Fio,
			},
			IsClosed: dbBlog.IsClose,
			Stats: &pb.Blog_Blog_Stats{
				ArticleCount:    dbBlog.TopicsCount,
				SubscriberCount: dbBlog.SubscriberCount,
			},
			LastArticle: &pb.Blog_LastArticle{
				Id:    dbBlog.LastTopicId,
				Title: dbBlog.LastTopicHead,
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
		var gender pb.Gender
		if dbBlogTopic.Sex == 0 {
			gender = pb.Gender_FEMALE
		} else {
			gender = pb.Gender_MALE
		}

		avatar := urlFormatter.GetAvatarUrl(dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

		article := &pb.Blog_Article{
			Id:    dbBlogTopic.TopicId,
			Title: dbBlogTopic.HeadTopic,
			Creation: &pb.Blog_Creation{
				User: &pb.Blog_UserLink{
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
