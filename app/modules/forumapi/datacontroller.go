package forumapi

import (
	"fantlab/db"
	"fantlab/pb"
	"fantlab/utils"
)

func getForumBlocks(dbForums []db.Forum, dbModerators map[uint32][]db.ForumModerator, urlFormatter utils.UrlFormatter) *pb.Forum_ForumBlocksResponse {
	var forumBlocks []*pb.Forum_ForumBlock

	currentForumBlockID := uint32(0) // f_forum_block.id начинаются с 1

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
					gender := utils.GetGender(dbModerator.UserID, dbModerator.Sex)
					avatar := urlFormatter.GetUserAvatarUrl(dbModerator.UserID, dbModerator.PhotoNumber)

					userLink := &pb.Common_UserLink{
						Id:     dbModerator.UserID,
						Login:  dbModerator.Login,
						Gender: gender,
						Avatar: avatar,
					}
					moderators = append(moderators, userLink)
				}

				gender := utils.GetGender(dbForum.UserID, dbForum.Sex)
				avatar := urlFormatter.GetUserAvatarUrl(dbForum.UserID, dbForum.PhotoNumber)

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
						Date: utils.ProtoTS(dbForum.LastMessageDate),
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

func getForumTopics(dbTopics []db.ForumTopic, urlFormatter utils.UrlFormatter) *pb.Forum_ForumTopicsResponse {
	//noinspection GoPreferNilSlice
	topics := []*pb.Forum_Topic{}

	for _, dbTopic := range dbTopics {
		var topicType pb.Forum_Topic_Type
		if dbTopic.TopicTypeID == 2 {
			topicType = pb.Forum_Topic_POLL
		} else {
			topicType = pb.Forum_Topic_TOPIC
		}

		creationUserGender := utils.GetGender(dbTopic.UserID, dbTopic.Sex)
		creationUserAvatar := urlFormatter.GetUserAvatarUrl(dbTopic.UserID, dbTopic.PhotoNumber)

		lastMessageUserGender := utils.GetGender(dbTopic.LastUserID, dbTopic.LastSex)
		lastMessageUserAvatar := urlFormatter.GetUserAvatarUrl(dbTopic.LastUserID, dbTopic.LastPhotoNumber)

		topic := &pb.Forum_Topic{
			Id:        dbTopic.TopicID,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbTopic.UserID,
					Login:  dbTopic.Login,
					Gender: creationUserGender,
					Avatar: creationUserAvatar,
				},
				Date: utils.ProtoTS(dbTopic.DateOfAdd),
			},
			IsClosed: dbTopic.IsClosed,
			IsPinned: dbTopic.IsPinned,
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
				Date: utils.ProtoTS(dbTopic.LastMessageDate),
			},
		}

		topics = append(topics, topic)
	}

	return &pb.Forum_ForumTopicsResponse{
		Topics: topics,
	}
}

func getTopic(shortTopic db.ShortForumTopic, dbMessages []db.ForumMessage, urlFormatter utils.UrlFormatter) *pb.Forum_TopicResponse {
	topic := &pb.Forum_Topic{
		Id:    shortTopic.TopicID,
		Title: shortTopic.TopicName,
	}

	forum := &pb.Forum_Forum{
		Id:    shortTopic.ForumID,
		Title: shortTopic.ForumName,
	}

	//noinspection GoPreferNilSlice
	messages := []*pb.Forum_TopicMessage{}

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText

		if dbMessage.IsCensored {
			text = ""
		}

		gender := utils.GetGender(dbMessage.UserID, dbMessage.Sex)
		avatar := urlFormatter.GetUserAvatarUrl(dbMessage.UserID, dbMessage.PhotoNumber)

		message := &pb.Forum_TopicMessage{
			Id: dbMessage.MessageID,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbMessage.UserID,
					Login:  dbMessage.Login,
					Gender: gender,
					Avatar: avatar,
					Class:  uint32(dbMessage.UserClass),
					Sign:   dbMessage.Sign,
				},
				Date: utils.ProtoTS(dbMessage.DateOfAdd),
			},
			Text:            text,
			IsCensored:      dbMessage.IsCensored,
			IsModerTagWorks: dbMessage.IsRed,
			Stats: &pb.Forum_TopicMessage_Stats{
				Rating: int32(dbMessage.VotePlus - dbMessage.VoteMinus),
			},
		}

		messages = append(messages, message)
	}

	return &pb.Forum_TopicResponse{
		Topic:    topic,
		Forum:    forum,
		Messages: messages,
	}
}
