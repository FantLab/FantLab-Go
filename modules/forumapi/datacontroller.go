package forumapi

import (
	"strconv"

	"fantlab/utils"
)

func getForumBlocks(dbForums []dbForum, dbModerators map[uint16][]dbModerator) forumBlocksWrapper {
	var forumBlocks []forumBlock

	currentForumBlockID := uint16(0) // f_forum_block.id начинаются с 1

	for _, dbForum := range dbForums {
		if dbForum.ForumBlockID != currentForumBlockID {
			forumBlock := forumBlock{
				ID:     dbForum.ForumBlockID,
				Title:  dbForum.ForumBlockName,
				Forums: []forum{},
			}
			forumBlocks = append(forumBlocks, forumBlock)
			currentForumBlockID = dbForum.ForumBlockID
		}
	}

	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockID == forumBlocks[index].ID {
				var moderators []userLink

				for _, dbModerator := range dbModerators[dbForum.ForumID] {
					userLink := userLink{
						ID:    dbModerator.UserID,
						Login: dbModerator.Login,
					}
					moderators = append(moderators, userLink)
				}

				forum := forum{
					ID:          dbForum.ForumID,
					Title:       dbForum.Name,
					Description: dbForum.Description,
					Moderators:  moderators,
					Stats: forumStats{
						TopicCount:   dbForum.TopicCount,
						MessageCount: dbForum.MessageCount,
					},
					LastMessage: lastMessage{
						ID: dbForum.LastMessageID,
						Topic: &topicLink{
							ID:    dbForum.LastTopicID,
							Title: dbForum.LastTopicName,
						},
						User: userLink{
							ID:    dbForum.LastUserID,
							Login: dbForum.LastUserName,
						},
						Date: utils.NewDateTime(dbForum.LastMessageDate),
					},
				}

				forumBlocks[index].Forums = append(forumBlocks[index].Forums, forum)

				break
			}
		}
	}

	return forumBlocksWrapper{forumBlocks}
}

func getForumTopics(dbTopics []dbForumTopic) forumTopicsWrapper {
	//noinspection GoPreferNilSlice
	topics := []forumTopic{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbTopic := range dbTopics {
		var topicType string
		if dbTopic.TopicTypeID == 2 {
			topicType = "poll"
		} else {
			topicType = "topic"
		}

		topic := forumTopic{
			ID:        dbTopic.TopicID,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: creation{
				User: userLink{
					ID:    dbTopic.UserID,
					Login: dbTopic.Login,
				},
				Date: utils.NewDateTime(dbTopic.DateOfAdd),
			},
			IsClosed: dbTopic.IsClosed,
			IsPinned: dbTopic.IsPinned,
			Stats: topicStats{
				MessageCount: dbTopic.MessageCount,
				ViewsCount:   dbTopic.Views,
			},
			LastMessage: lastMessage{
				ID: dbTopic.LastMessageID,
				User: userLink{
					ID:    dbTopic.LastUserID,
					Login: dbTopic.LastUserName,
				},
				Date: utils.NewDateTime(dbTopic.LastMessageDate),
			},
		}

		topics = append(topics, topic)
	}

	return forumTopicsWrapper{topics}
}

func getTopicMessages(dbMessages []dbForumMessage, imageUrl string) topicMessagesWrapper {
	//noinspection GoPreferNilSlice
	messages := []topicMessage{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText

		if dbMessage.IsCensored {
			text = ""
		}

		var gender string
		if dbMessage.Sex == 0 {
			gender = "f"
		} else {
			gender = "m"
		}

		var avatar string
		if dbMessage.PhotoNumber != 0 {
			userId := strconv.FormatUint(uint64(dbMessage.UserID), 10)
			photoNumber := strconv.FormatUint(uint64(dbMessage.PhotoNumber), 10)
			avatar = imageUrl + "/users/" + userId + "_" + photoNumber
		}

		message := topicMessage{
			ID: dbMessage.MessageID,
			Creation: creation{
				User: userLink{
					ID:     dbMessage.UserID,
					Login:  dbMessage.Login,
					Gender: gender,
					Avatar: avatar,
					Class:  dbMessage.UserClass,
					Sign:   dbMessage.Sign,
				},
				Date: utils.NewDateTime(dbMessage.DateOfAdd),
			},
			Text:            text,
			IsCensored:      dbMessage.IsCensored,
			IsModerTagWorks: dbMessage.IsRed,
			Stats: messageStats{
				PlusCount:  dbMessage.VotePlus,
				MinusCount: dbMessage.VoteMinus,
			},
		}

		messages = append(messages, message)
	}

	return topicMessagesWrapper{messages}
}
