package forumapi

import (
	"fantlab/protobuf/generated/fantlab/apimodels"
	"fantlab/utils"
)

func getForumBlocks(dbForums []dbForum, dbModerators map[uint32][]dbModerator) *apimodels.Forum_ForumBlocksResponse {
	var forumBlocks []*apimodels.Forum_ForumBlock

	currentForumBlockID := uint32(0) // f_forum_block.id начинаются с 1

	for _, dbForum := range dbForums {
		if dbForum.ForumBlockID != currentForumBlockID {
			forumBlock := apimodels.Forum_ForumBlock{
				Id:     dbForum.ForumBlockID,
				Title:  dbForum.ForumBlockName,
				Forums: []*apimodels.Forum_Forum{},
			}
			forumBlocks = append(forumBlocks, &forumBlock)
			currentForumBlockID = dbForum.ForumBlockID
		}
	}

	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockID == forumBlocks[index].GetId() {
				var moderators []*apimodels.Forum_UserLink

				for _, dbModerator := range dbModerators[dbForum.ForumID] {
					userLink := &apimodels.Forum_UserLink{
						Id:    dbModerator.UserID,
						Login: dbModerator.Login,
					}
					moderators = append(moderators, userLink)
				}

				forum := apimodels.Forum_Forum{
					Id:               dbForum.ForumID,
					Title:            dbForum.Name,
					ForumDescription: dbForum.Description,
					Users:            moderators,
					Stats: &apimodels.Forum_Forum_Stats{
						TopicCount:   dbForum.TopicCount,
						MessageCount: dbForum.MessageCount,
					},
					LastMessage: &apimodels.Forum_LastMessage{
						Id: dbForum.LastMessageID,
						Topic: &apimodels.Forum_TopicLink{
							Id:    dbForum.LastTopicID,
							Title: dbForum.LastTopicName,
						},
						User: &apimodels.Forum_UserLink{
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

	return &apimodels.Forum_ForumBlocksResponse{
		ForumBlocks: forumBlocks,
	}
}

func getForumTopics(dbTopics []dbForumTopic) *apimodels.Forum_ForumTopicsResponse {
	//noinspection GoPreferNilSlice
	topics := []*apimodels.Forum_Topic{}

	for _, dbTopic := range dbTopics {
		var topicType apimodels.Forum_Topic_Type
		if dbTopic.TopicTypeID == 2 {
			topicType = apimodels.Forum_Topic_POLL
		} else {
			topicType = apimodels.Forum_Topic_TOPIC
		}

		topic := &apimodels.Forum_Topic{
			Id:        dbTopic.TopicID,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: &apimodels.Forum_Creation{
				User: &apimodels.Forum_UserLink{
					Id:    dbTopic.UserID,
					Login: dbTopic.Login,
				},
				Date: utils.ProtoTS(dbTopic.DateOfAdd),
			},
			IsClosed: dbTopic.IsClosed,
			IsPinned: dbTopic.IsPinned,
			Stats: &apimodels.Forum_Topic_Stats{
				MessageCount: dbTopic.MessageCount,
				ViewsCount:   dbTopic.Views,
			},
			LastMessage: &apimodels.Forum_LastMessage{
				Id: dbTopic.LastMessageID,
				User: &apimodels.Forum_UserLink{
					Id:    dbTopic.LastUserID,
					Login: dbTopic.LastUserName,
				},
				Date: utils.ProtoTS(dbTopic.LastMessageDate),
			},
		}

		topics = append(topics, topic)
	}

	return &apimodels.Forum_ForumTopicsResponse{
		Topics: topics,
	}
}

func getTopicMessages(dbMessages []dbForumMessage, urlFormatter utils.UrlFormatter) *apimodels.Forum_TopicMessagesResponse {
	//noinspection GoPreferNilSlice
	messages := []*apimodels.Forum_TopicMessage{}

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText

		if dbMessage.IsCensored {
			text = ""
		}

		var gender apimodels.Gender
		if dbMessage.Sex == 0 {
			gender = apimodels.Gender_FEMALE
		} else {
			gender = apimodels.Gender_MALE
		}

		avatar := urlFormatter.GetAvatarUrl(dbMessage.UserID, dbMessage.PhotoNumber)

		message := &apimodels.Forum_TopicMessage{
			Id: dbMessage.MessageID,
			Creation: &apimodels.Forum_Creation{
				User: &apimodels.Forum_UserLink{
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
			Stats: &apimodels.Forum_TopicMessage_Stats{
				PlusCount:  dbMessage.VotePlus,
				MinusCount: dbMessage.VoteMinus,
			},
		}

		messages = append(messages, message)
	}

	return &apimodels.Forum_TopicMessagesResponse{
		Messages: messages,
	}
}
