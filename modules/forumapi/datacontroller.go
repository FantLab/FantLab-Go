package forumapi

import (
	"fantlab/protobuf/generated/fantlab/pb"
	"fantlab/utils"
)

func getForumBlocks(dbForums []dbForum, dbModerators map[uint32][]dbModerator) *pb.Forum_ForumBlocksResponse {
	var forumBlocks []*pb.Forum_ForumBlock

	currentForumBlockID := uint32(0) // f_forum_block.id начинаются с 1

	for _, dbForum := range dbForums {
		if dbForum.ForumBlockID != currentForumBlockID {
			forumBlock := pb.Forum_ForumBlock{
				XXX_Id: dbForum.ForumBlockID,
				Title:  dbForum.ForumBlockName,
				Forums: []*pb.Forum_Forum{},
			}
			forumBlocks = append(forumBlocks, &forumBlock)
			currentForumBlockID = dbForum.ForumBlockID
		}
	}

	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockID == forumBlocks[index].GetXXX_Id() {
				var moderators []*pb.Forum_UserLink

				for _, dbModerator := range dbModerators[dbForum.ForumID] {
					userLink := &pb.Forum_UserLink{
						Id:    dbModerator.UserID,
						Login: dbModerator.Login,
					}
					moderators = append(moderators, userLink)
				}

				forum := pb.Forum_Forum{
					Id:          dbForum.ForumID,
					Title:       dbForum.Name,
					Description: dbForum.Description,
					Moderators:  moderators,
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
						User: &pb.Forum_UserLink{
							Id:    dbForum.LastUserID,
							Login: dbForum.LastUserName,
						},
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

func getForumTopics(dbTopics []dbForumTopic) *pb.Forum_ForumTopicsResponse {
	//noinspection GoPreferNilSlice
	topics := []*pb.Forum_Topic{}

	for _, dbTopic := range dbTopics {
		var topicType pb.Forum_Topic_Type
		if dbTopic.TopicTypeID == 2 {
			topicType = pb.Forum_Topic_poll
		} else {
			topicType = pb.Forum_Topic_topic
		}

		topic := &pb.Forum_Topic{
			Id:        dbTopic.TopicID,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: &pb.Forum_Creation{
				User: &pb.Forum_UserLink{
					Id:    dbTopic.UserID,
					Login: dbTopic.Login,
				},
				Date: utils.ProtoTS(dbTopic.DateOfAdd),
			},
			IsClosed: dbTopic.IsClosed,
			IsPinned: dbTopic.IsPinned,
			Stats: &pb.Forum_Topic_Stats{
				MessageCount: dbTopic.MessageCount,
				ViewsCount:   dbTopic.Views,
			},
			LastMessage: &pb.Forum_LastMessage{
				Id: dbTopic.LastMessageID,
				User: &pb.Forum_UserLink{
					Id:    dbTopic.LastUserID,
					Login: dbTopic.LastUserName,
				},
				Date: utils.ProtoTS(dbTopic.LastMessageDate),
			},
		}

		topics = append(topics, topic)
	}

	return &pb.Forum_ForumTopicsResponse{
		Topics: topics,
	}
}

func getTopicMessages(dbMessages []dbForumMessage, urlFormatter utils.UrlFormatter) *pb.Forum_TopicMessagesResponse {
	//noinspection GoPreferNilSlice
	messages := []*pb.Forum_TopicMessage{}

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText

		if dbMessage.IsCensored {
			text = ""
		}

		var gender pb.Gender
		if dbMessage.Sex == 0 {
			gender = pb.Gender_female
		} else {
			gender = pb.Gender_male
		}

		avatar := urlFormatter.GetAvatarUrl(dbMessage.UserID, dbMessage.PhotoNumber)

		message := &pb.Forum_TopicMessage{
			Id: dbMessage.MessageID,
			Creation: &pb.Forum_Creation{
				User: &pb.Forum_UserLink{
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
				PlusCount:  dbMessage.VotePlus,
				MinusCount: dbMessage.VoteMinus,
			},
		}

		messages = append(messages, message)
	}

	return &pb.Forum_TopicMessagesResponse{
		Messages: messages,
	}
}
