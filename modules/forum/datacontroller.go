package forumapi

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
						Topic: topicLink{
							ID:    dbForum.LastTopicID,
							Title: dbForum.LastTopicName,
						},
						User: userLink{
							ID:    dbForum.LastUserID,
							Login: dbForum.LastUserName,
						},
						Date: dbForum.LastMessageDate.Unix(),
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
		topic := forumTopic{
			ID:        dbTopic.TopicID,
			Title:     dbTopic.Name,
			ForumType: dbTopic.TopicTypeID,
			Creation: creation{
				User: userLink{
					ID:    dbTopic.UserID,
					Login: dbTopic.Login,
				},
				Date: dbTopic.DateOfAdd.Unix(),
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
				Date: dbTopic.LastMessageDate.Unix(),
			},
		}

		topics = append(topics, topic)
	}

	return forumTopicsWrapper{topics}
}

func getTopicMessages(dbMessages []dbForumMessage) topicMessagesWrapper {
	//noinspection GoPreferNilSlice
	messages := []topicMessage{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText

		if dbMessage.IsCensored {
			text = ""
		}

		message := topicMessage{
			ID: dbMessage.MessageID,
			Creation: creation{
				User: userLink{
					ID:          dbMessage.UserID,
					Login:       dbMessage.Login,
					Gender:      dbMessage.Sex,
					PhotoNumber: dbMessage.PhotoNumber,
					Class:       dbMessage.UserClass,
					Sign:        dbMessage.Sign,
				},
				Date: dbMessage.DateOfAdd.Unix(),
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
