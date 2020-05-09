package converters

import (
	"fantlab/base/protobuf/pbutils"
	"fantlab/core/config"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
)

func GetForumBlocks(dbForums []db.Forum, dbModerators map[uint64][]db.ForumModerator, cfg *config.AppConfig) *pb.Forum_ForumBlocksResponse {
	var forumBlocks []*pb.Forum_ForumBlock

	currentForumBlockID := uint64(0) // f_forum_block.id начинаются с 1

	for _, dbForum := range dbForums {
		if dbForum.ForumBlockID != currentForumBlockID {
			forumBlock := pb.Forum_ForumBlock{
				Id:     dbForum.ForumBlockID,
				Title:  dbForum.ForumBlockName,
				Forums: []*pb.Forum_Forum{},
			}
			forumBlocks = append(forumBlocks, &forumBlock)
			currentForumBlockID = dbForum.ForumBlockID
		}
	}

	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockID == forumBlocks[index].GetId() {
				var moderators []*pb.Common_UserLink

				for _, dbModerator := range dbModerators[dbForum.ForumID] {
					gender := helpers.GetGender(dbModerator.UserID, dbModerator.Sex)
					avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbModerator.UserID, dbModerator.PhotoNumber)

					userLink := &pb.Common_UserLink{
						Id:     dbModerator.UserID,
						Login:  dbModerator.Login,
						Gender: gender,
						Avatar: avatar,
					}
					moderators = append(moderators, userLink)
				}

				gender := helpers.GetGender(dbForum.UserID, dbForum.Sex)
				avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbForum.UserID, dbForum.PhotoNumber)

				forum := pb.Forum_Forum{
					Id:               dbForum.ForumID,
					Title:            dbForum.Name,
					ForumDescription: dbForum.Description,
					Moderators:       moderators,
					Stats: &pb.Forum_Forum_Stats{
						TopicCount:   dbForum.TopicCount,
						MessageCount: dbForum.MessageCount,
					},
					LastMessage: &pb.Forum_LastMessage{
						Id: dbForum.LastMessageID,
						Topic: &pb.Forum_TopicLink{
							Id:    dbForum.LastTopicID,
							Title: dbForum.LastTopicName,
						},
						User: &pb.Common_UserLink{
							Id:     dbForum.UserID,
							Login:  dbForum.Login,
							Gender: gender,
							Avatar: avatar,
						},
						Text: dbForum.LastMessageText,
						Date: pbutils.TimestampProto(dbForum.LastMessageDate),
					},
				}

				forumBlocks[index].Forums = append(forumBlocks[index].Forums, &forum)

				break
			}
		}
	}

	return &pb.Forum_ForumBlocksResponse{
		ForumBlocks: forumBlocks,
	}
}

func GetForumTopics(dbResponse *db.ForumTopicsDBResponse, page, limit uint64, cfg *config.AppConfig) *pb.Forum_ForumTopicsResponse {
	//noinspection GoPreferNilSlice
	topics := []*pb.Forum_Topic{}

	for _, dbTopic := range dbResponse.Topics {
		var topicType pb.Forum_Topic_Type
		if dbTopic.TopicTypeID == 2 {
			topicType = pb.Forum_Topic_POLL
		} else {
			topicType = pb.Forum_Topic_TOPIC
		}

		creationUserGender := helpers.GetGender(dbTopic.UserID, dbTopic.Sex)
		creationUserAvatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbTopic.UserID, dbTopic.PhotoNumber)

		lastMessageUserGender := helpers.GetGender(dbTopic.LastUserID, dbTopic.LastSex)
		lastMessageUserAvatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbTopic.LastUserID, dbTopic.LastPhotoNumber)

		topic := &pb.Forum_Topic{
			Id:        dbTopic.TopicId,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbTopic.UserID,
					Login:  dbTopic.Login,
					Gender: creationUserGender,
					Avatar: creationUserAvatar,
				},
				Date: pbutils.TimestampProto(dbTopic.DateOfAdd),
			},
			IsClosed: dbTopic.IsClosed != 0,
			IsPinned: dbTopic.IsPinned != 0,
			Stats: &pb.Forum_Topic_Stats{
				MessageCount: dbTopic.MessageCount,
				ViewCount:    dbTopic.Views,
			},
			LastMessage: &pb.Forum_LastMessage{
				Id: dbTopic.LastMessageID,
				User: &pb.Common_UserLink{
					Id:     dbTopic.LastUserID,
					Login:  dbTopic.LastLogin,
					Gender: lastMessageUserGender,
					Avatar: lastMessageUserAvatar,
				},
				Text: dbTopic.LastMessageText,
				Date: pbutils.TimestampProto(dbTopic.LastMessageDate),
			},
		}

		topics = append(topics, topic)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalTopicsCount, limit)

	return &pb.Forum_ForumTopicsResponse{
		Topics: topics,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetTopic(dbResponse *db.ForumTopicMessagesDBResponse, page, limit uint64, cfg *config.AppConfig) *pb.Forum_ForumTopicResponse {
	topic := &pb.Forum_Topic{
		Id:    dbResponse.Topic.TopicID,
		Title: dbResponse.Topic.TopicName,
	}

	forum := &pb.Forum_Forum{
		Id:    dbResponse.Topic.ForumID,
		Title: dbResponse.Topic.ForumName,
	}

	var pinnedMessage *pb.Forum_TopicMessage

	if dbResponse.PinnedFirstMessage != (db.ForumMessage{}) {
		pinnedMessage = convertMessage(&dbResponse.PinnedFirstMessage, cfg)
	}

	//noinspection GoPreferNilSlice
	messages := []*pb.Forum_TopicMessage{}

	for _, dbMessage := range dbResponse.Messages {
		messages = append(messages, convertMessage(&dbMessage, cfg))
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalMessagesCount, limit)

	return &pb.Forum_ForumTopicResponse{
		Topic:         topic,
		Forum:         forum,
		PinnedMessage: pinnedMessage,
		Messages:      messages,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetForumTopicMessage(dbMessage *db.ForumMessage, cfg *config.AppConfig) *pb.Forum_ForumMessageResponse {
	message := convertMessage(dbMessage, cfg)

	return &pb.Forum_ForumMessageResponse{
		Message: message,
	}
}

func convertMessage(dbMessage *db.ForumMessage, cfg *config.AppConfig) *pb.Forum_TopicMessage {
	text := dbMessage.MessageText

	if dbMessage.IsCensored != 0 {
		text = cfg.CensorshipText
	}

	gender := helpers.GetGender(dbMessage.UserID, dbMessage.Sex)
	avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbMessage.UserID, dbMessage.PhotoNumber)

	return &pb.Forum_TopicMessage{
		Id: dbMessage.MessageID,
		Creation: &pb.Common_Creation{
			User: &pb.Common_UserLink{
				Id:     dbMessage.UserID,
				Login:  dbMessage.Login,
				Gender: gender,
				Avatar: avatar,
				Class:  helpers.UserClassMap[dbMessage.UserClass],
				Sign:   dbMessage.Sign,
			},
			Date: pbutils.TimestampProto(dbMessage.DateOfAdd),
		},
		Text:       text,
		IsCensored: dbMessage.IsCensored == 1,
		Stats: &pb.Forum_TopicMessage_Stats{
			Rating: int64(dbMessage.VotePlus - dbMessage.VoteMinus),
		},
	}
}

func GetForumTopicMessageDraft(dbMessageDraft *db.ForumMessageDraft, cfg *config.AppConfig) *pb.Forum_ForumMessageDraftResponse {
	gender := helpers.GetGender(dbMessageDraft.UserID, dbMessageDraft.Sex)
	avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbMessageDraft.UserID, dbMessageDraft.PhotoNumber)

	return &pb.Forum_ForumMessageDraftResponse{
		MessageDraft: &pb.Forum_TopicMessageDraft{
			TopicId: dbMessageDraft.TopicId,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbMessageDraft.UserID,
					Login:  dbMessageDraft.Login,
					Gender: gender,
					Avatar: avatar,
					Class:  helpers.UserClassMap[dbMessageDraft.UserClass],
					Sign:   dbMessageDraft.Sign,
				},
				Date: pbutils.TimestampProto(dbMessageDraft.DateOfAdd),
			},
			Text: dbMessageDraft.Message,
		},
	}
}
