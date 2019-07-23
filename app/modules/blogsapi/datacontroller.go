package blogsapi

import (
	"fantlab/db"
	"fantlab/pb"
	"fantlab/shared"
	"fantlab/utils"
)

func getCommunities(dbCommunities []db.Community, cfg *shared.AppConfig) *pb.Blog_CommunitiesResponse {
	var mainCommunities []*pb.Blog_Community
	var additionalCommunities []*pb.Blog_Community

	for _, dbCommunity := range dbCommunities {
		userGender := utils.GetGender(dbCommunity.LastUserId, dbCommunity.LastSex)
		userAvatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbCommunity.LastUserId, dbCommunity.LastPhotoNumber)
		communityAvatar := utils.GetCommunityAvatarUrl(cfg.ImagesBaseURL, dbCommunity.BlogId)

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

func getCommunity(dbResponse *db.CommunityDBResponse, page, limit uint32, cfg *shared.AppConfig) *pb.Blog_CommunityResponse {
	communityAvatar := utils.GetCommunityAvatarUrl(cfg.ImagesBaseURL, dbResponse.Community.BlogId)

	community := &pb.Blog_Community{
		Id:     dbResponse.Community.BlogId,
		Title:  dbResponse.Community.Name,
		Rules:  dbResponse.Community.Rules,
		Avatar: communityAvatar,
	}

	var moderators []*pb.Common_UserLink

	for _, dbModerator := range dbResponse.Moderators {
		gender := utils.GetGender(dbModerator.UserID, dbModerator.Sex)
		avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbModerator.UserID, dbModerator.PhotoNumber)

		moderator := &pb.Common_UserLink{
			Id:     dbModerator.UserID,
			Login:  dbModerator.Login,
			Gender: gender,
			Avatar: avatar,
		}

		moderators = append(moderators, moderator)
	}

	var authors []*pb.Common_UserLink

	for _, dbAuthor := range dbResponse.Authors {
		gender := utils.GetGender(dbAuthor.UserID, dbAuthor.Sex)
		avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbAuthor.UserID, dbAuthor.PhotoNumber)

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

	for _, dbTopic := range dbResponse.Topics {
		gender := utils.GetGender(dbTopic.UserId, dbTopic.Sex)
		avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbTopic.UserId, uint32(dbTopic.PhotoNumber))

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

	pageCount := utils.GetPageCount(dbResponse.TotalTopicsCount, limit)

	return &pb.Blog_CommunityResponse{
		Community:  community,
		Moderators: moderators,
		Authors:    authors,
		Articles:   articles,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func getBlogs(dbResponse *db.BlogsDBResponse, page, limit uint32, cfg *shared.AppConfig) *pb.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*pb.Blog_Blog{}

	for _, dbBlog := range dbResponse.Blogs {
		gender := utils.GetGender(dbBlog.UserId, dbBlog.Sex)
		avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlog.UserId, dbBlog.PhotoNumber)

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

	pageCount := utils.GetPageCount(dbResponse.TotalCount, limit)

	return &pb.Blog_BlogsResponse{
		Blogs: blogs,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func getBlog(dbResponse *db.BlogDBResponse, page, limit uint32, cfg *shared.AppConfig) *pb.Blog_BlogResponse {
	//noinspection GoPreferNilSlice
	var articles = []*pb.Blog_Article{}

	for _, dbBlogTopic := range dbResponse.Topics {
		gender := utils.GetGender(dbBlogTopic.UserId, dbBlogTopic.Sex)
		avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

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

	pageCount := utils.GetPageCount(dbResponse.TotalTopicsCount, limit)

	return &pb.Blog_BlogResponse{
		Articles: articles,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func getArticle(dbBlogTopic *db.BlogTopic, cfg *shared.AppConfig) *pb.Blog_BlogArticleResponse {
	gender := utils.GetGender(dbBlogTopic.UserId, dbBlogTopic.Sex)
	avatar := utils.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlogTopic.UserId, uint32(dbBlogTopic.PhotoNumber))

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

	return &pb.Blog_BlogArticleResponse{
		Article: article,
	}
}
