package converters

import (
	"fantlab/base/protobuf/pbutils"
	"fantlab/core/config"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
)

func GetCommunities(dbCommunities []db.Community, cfg *config.AppConfig) *pb.Blog_CommunitiesResponse {
	var mainCommunities []*pb.Blog_Community
	var additionalCommunities []*pb.Blog_Community

	for _, dbCommunity := range dbCommunities {
		userGender := helpers.GetGender(dbCommunity.LastUserId, dbCommunity.LastSex)
		userAvatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbCommunity.LastUserId, dbCommunity.LastPhotoNumber)
		communityAvatar := helpers.GetCommunityAvatarUrl(cfg.ImagesBaseURL, dbCommunity.BlogId)

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
				Date: pbutils.TimestampProto(dbCommunity.LastTopicDate),
			},
		}

		if dbCommunity.IsPublic != 0 {
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

func GetCommunity(dbResponse *db.CommunityTopicsDBResponse, page, limit uint64, cfg *config.AppConfig) *pb.Blog_CommunityResponse {
	communityAvatar := helpers.GetCommunityAvatarUrl(cfg.ImagesBaseURL, dbResponse.Community.BlogId)

	community := &pb.Blog_Community{
		Id:     dbResponse.Community.BlogId,
		Title:  dbResponse.Community.Name,
		Rules:  dbResponse.Community.Rules,
		Avatar: communityAvatar,
	}

	var moderators []*pb.Common_UserLink

	for _, dbModerator := range dbResponse.Moderators {
		gender := helpers.GetGender(dbModerator.UserID, dbModerator.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbModerator.UserID, dbModerator.PhotoNumber)

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
		gender := helpers.GetGender(dbAuthor.UserID, dbAuthor.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbAuthor.UserID, dbAuthor.PhotoNumber)

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
		gender := helpers.GetGender(dbTopic.UserId, dbTopic.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbTopic.UserId, dbTopic.PhotoNumber)

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
				Date: pbutils.TimestampProto(dbTopic.DateOfAdd),
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

	pageCount := helpers.CalculatePageCount(dbResponse.TotalTopicsCount, limit)

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

func GetBlogs(dbResponse *db.BlogsDBResponse, page, limit uint64, cfg *config.AppConfig) *pb.Blog_BlogsResponse {
	//noinspection GoPreferNilSlice
	var blogs = []*pb.Blog_Blog{}

	for _, dbBlog := range dbResponse.Blogs {
		gender := helpers.GetGender(dbBlog.UserId, dbBlog.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlog.UserId, dbBlog.PhotoNumber)

		blog := &pb.Blog_Blog{
			Id: dbBlog.BlogId,
			User: &pb.Common_UserLink{
				Id:     dbBlog.UserId,
				Login:  dbBlog.Login,
				Name:   dbBlog.Fio,
				Gender: gender,
				Avatar: avatar,
			},
			IsClosed: dbBlog.IsClose != 0,
			Stats: &pb.Blog_Blog_Stats{
				ArticleCount:    dbBlog.TopicsCount,
				SubscriberCount: dbBlog.SubscriberCount,
			},
			LastArticle: &pb.Blog_LastArticle{
				Id:    dbBlog.LastTopicId,
				Title: dbBlog.LastTopicHead,
				Date:  pbutils.TimestampProto(dbBlog.LastTopicDate),
			},
		}

		blogs = append(blogs, blog)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalCount, limit)

	return &pb.Blog_BlogsResponse{
		Blogs: blogs,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetBlog(dbResponse *db.BlogTopicsDBResponse, viewCounts []uint64, page, limit uint64, cfg *config.AppConfig) *pb.Blog_BlogResponse {
	//noinspection GoPreferNilSlice
	var articles = []*pb.Blog_Article{}

	for index, dbBlogTopic := range dbResponse.Topics {
		gender := helpers.GetGender(dbBlogTopic.UserId, dbBlogTopic.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlogTopic.UserId, dbBlogTopic.PhotoNumber)

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
				Date: pbutils.TimestampProto(dbBlogTopic.DateOfAdd),
			},
			Tags: dbBlogTopic.Tags,
			Stats: &pb.Blog_Article_Stats{
				LikeCount:    dbBlogTopic.LikesCount,
				ViewCount:    viewCounts[index],
				CommentCount: dbBlogTopic.CommentsCount,
			},
		}

		articles = append(articles, article)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalTopicsCount, limit)

	return &pb.Blog_BlogResponse{
		Articles: articles,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetArticle(dbBlogTopic *db.BlogTopic, viewCount uint64, attachments []helpers.File, cfg *config.AppConfig) *pb.Blog_BlogArticleResponse {
	gender := helpers.GetGender(dbBlogTopic.UserId, dbBlogTopic.Sex)
	avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbBlogTopic.UserId, dbBlogTopic.PhotoNumber)

	//noinspection GoPreferNilSlice
	attaches := []*pb.Common_Attachment{}
	for _, attachment := range attachments {
		attaches = append(attaches, &pb.Common_Attachment{
			Name: attachment.Name,
			Size: attachment.Size,
		})
	}

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
			Date: pbutils.TimestampProto(dbBlogTopic.DateOfAdd),
		},
		Text:        dbBlogTopic.MessageText,
		Tags:        dbBlogTopic.Tags,
		Attachments: attaches,
		Stats: &pb.Blog_Article_Stats{
			LikeCount:    dbBlogTopic.LikesCount,
			ViewCount:    viewCount,
			CommentCount: dbBlogTopic.CommentsCount,
		},
	}

	return &pb.Blog_BlogArticleResponse{
		Article: article,
	}
}

func GetBlogArticleComments(entries []db.BlogTopicComment, totalCount uint64, cfg *config.AppConfig) *pb.Blog_BlogArticleCommentsResponse {
	messageTable := make(map[uint64]*pb.Blog_Comment)

	for _, entry := range entries {
		messageTable[entry.MessageId] = &pb.Blog_Comment{
			Id: entry.MessageId,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     entry.UserId,
					Login:  entry.UserLogin,
					Gender: helpers.GetGender(entry.UserId, entry.UserSex),
					Avatar: helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, entry.UserId, entry.UserPhotoNumber),
				},
				Date: pbutils.TimestampProto(entry.DateOfAdd),
			},
			Text:       entry.Text,
			IsCensored: entry.IsCensored > 0,
		}
	}

	response := new(pb.Blog_BlogArticleCommentsResponse)

	for _, entry := range entries {
		comment := messageTable[entry.MessageId]
		if comment == nil {
			continue
		}

		parentComment := messageTable[entry.ParentMessageId]

		if parentComment == nil {
			response.Comments = append(response.Comments, comment)
		} else {
			parentComment.Answers = append(parentComment.Answers, comment)
		}
	}

	response.TotalCount = totalCount

	return response
}

func GetBlogArticleComment(dbComment *db.BlogTopicComment, cfg *config.AppConfig) *pb.Blog_BlogArticleCommentResponse {
	comment := &pb.Blog_Comment{
		Id: dbComment.MessageId,
		Creation: &pb.Common_Creation{
			User: &pb.Common_UserLink{
				Id:     dbComment.UserId,
				Login:  dbComment.UserLogin,
				Gender: helpers.GetGender(dbComment.UserId, dbComment.UserSex),
				Avatar: helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbComment.UserId, dbComment.UserPhotoNumber),
			},
			Date: pbutils.TimestampProto(dbComment.DateOfAdd),
		},
		Text:       dbComment.Text,
		IsCensored: dbComment.IsCensored > 0,
	}

	return &pb.Blog_BlogArticleCommentResponse{
		Comment: comment,
	}
}
