package blogsapi

import (
	"fantlab/pb"
	"fantlab/utils"
)

func getCommunities(dbCommunities []dbCommunity, urlFormatter utils.UrlFormatter) *pb.Blog_CommunitiesResponse {
	var mainCommunities []*pb.Blog_Community
	var additionalCommunities []*pb.Blog_Community

	for _, dbCommunity := range dbCommunities {
		userGender := utils.GetGender(dbCommunity.LastSex)
		userAvatar := urlFormatter.GetUserAvatarUrl(dbCommunity.LastUserId, dbCommunity.LastPhotoNumber)
		communityAvatar := urlFormatter.GetCommunityAvatarUrl(dbCommunity.BlogId)

		community := &pb.Blog_Community{
			Id:                   dbCommunity.BlogId,
			Title:                dbCommunity.Name,
			CommunityDescription: dbCommunity.Description,
			Avatar:               communityAvatar,
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
					Gender: userGender,
					Avatar: userAvatar,
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

func getCommunity(dbCommunity dbCommunity,
	dbModerators []dbModerator,
	dbAuthors []dbAuthor,
	dbTopics []dbTopic,
	urlFormatter utils.UrlFormatter) *pb.Blog_CommunityResponse {
	communityAvatar := urlFormatter.GetCommunityAvatarUrl(dbCommunity.BlogId)

	community := &pb.Blog_Community{
		Id:     dbCommunity.BlogId,
		Title:  dbCommunity.Name,
		Rules:  dbCommunity.Rules,
		Avatar: communityAvatar,
	}

	var moderators []*pb.Common_UserLink

	for _, dbModerator := range dbModerators {
		gender := utils.GetGender(dbModerator.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbModerator.UserID, dbModerator.PhotoNumber)

		moderator := &pb.Common_UserLink{
			Id:     dbModerator.UserID,
			Login:  dbModerator.Login,
			Gender: gender,
			Avatar: avatar,
		}

		moderators = append(moderators, moderator)
	}

	var authors []*pb.Common_UserLink

	for _, dbAuthor := range dbAuthors {
		gender := utils.GetGender(dbAuthor.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbAuthor.UserID, dbAuthor.PhotoNumber)

		author := &pb.Common_UserLink{
			Id:     dbAuthor.UserID,
			Login:  dbAuthor.Login,
			Gender: gender,
			Avatar: avatar,
		}

		authors = append(authors, author)
	}

	//noinspection GoPreferNilSlice
	articles := []*pb.Blog_Article{}

	for _, dbTopic := range dbTopics {
		gender := utils.GetGender(dbTopic.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbTopic.UserId, uint32(dbTopic.PhotoNumber))

		article := &pb.Blog_Article{
			Id:    dbTopic.TopicId,
			Title: dbTopic.HeadTopic,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbTopic.UserId,
					Login:  dbTopic.Login,
					Gender: gender,
					Avatar: avatar,
				},
				Date: utils.ProtoTS(dbTopic.DateOfAdd),
			},
			Text: dbTopic.MessageText,
			Tags: dbTopic.Tags,
			Stats: &pb.Blog_Article_Stats{
				LikeCount:    dbTopic.LikesCount,
				ViewCount:    dbTopic.Views,
				CommentCount: dbTopic.CommentsCount,
			},
		}

		articles = append(articles, article)
	}

	return &pb.Blog_CommunityResponse{
		Community:  community,
		Moderators: moderators,
		Authors:    authors,
		Articles:   articles,
	}
}

func getBlogs(dbBlogs []dbBlog, urlFormatter utils.UrlFormatter) *pb.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*pb.Blog_Blog{}

	for _, dbBlog := range dbBlogs {
		gender := utils.GetGender(dbBlog.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbBlog.UserId, dbBlog.PhotoNumber)

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
				Date:  utils.ProtoTS(dbBlog.LastTopicDate),
			},
		}

		blogs = append(blogs, blog)
	}

	return &pb.Blog_BlogsResponse{
		Blogs: blogs,
	}
}

func getBlog(dbBlogTopics []dbTopic, urlFormatter utils.UrlFormatter) *pb.Blog_BlogResponse {
	//noinspection GoPreferNilSlice
	var articles = []*pb.Blog_Article{}

	for _, dbBlogTopic := range dbBlogTopics {
		gender := utils.GetGender(dbBlogTopic.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

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

	return &pb.Blog_BlogResponse{
		Articles: articles,
	}
}
