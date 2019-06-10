package main

// Блок форумов
type ForumBlock struct {
	Id     uint16  `json:"-"`
	Title  string  `json:"block_title"`
	Forums []Forum `json:"forums"`
}

// Форум
type Forum struct {
	Id          uint16      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Moderators  []userLink  `json:"moderators"`
	Stats       forumStats  `json:"stats"`
	LastMessage lastMessage `json:"last_message"`
}

// Тема
type ForumTopic struct {
	Id          uint32      `json:"id"`
	Title       string      `json:"title"`
	Type        uint16      `json:"type"`
	Creation    creation    `json:"creation"`
	IsClosed    bool        `json:"is_closed"`
	IsPinned    bool        `json:"is_pinned"`
	Stats       topicStats  `json:"stats"`
	LastMessage lastMessage `json:"last_message"`
}

// Сообщение в форуме
// IsCensored - сообщение изъято модератором
// IsModerTagWorks - рендерить ли в сообщении содержимое тега [moder] (true, только если тег добавлен модератором)
type TopicMessage struct {
	Id              uint32       `json:"id"`
	Creation        creation     `json:"creation"`
	Text            string       `json:"text"`
	IsCensored      bool         `json:"is_censored"`
	IsModerTagWorks bool         `json:"is_moder_tag_works"`
	Stats           messageStats `json:"stats"`
}

// Статистика форума
type forumStats struct {
	TopicCount   uint32 `json:"topic_count"`
	MessageCount uint32 `json:"message_count"`
}

// Последнее сообщение в форуме
type lastMessage struct {
	Id    uint32    `json:"id"`
	Topic topicLink `json:"topic,omitempty"`
	User  userLink  `json:"user"`
	Date  int64     `json:"date"`
}

// Ссылка на тему форума
type topicLink struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

// Ссылка на пользователя
// PhotoNumber - порядковый номер фото (https://data.fantlab.ru/images/users/{UserId}_{PhotoNumber}). Если 0 - его нет.
type userLink struct {
	Id          uint32 `json:"id"`
	Login       string `json:"login"`
	Gender      uint8  `json:"gender,omitempty"`
	PhotoNumber uint16 `json:"photo_number,omitempty"`
	Class       uint8  `json:"class,omitempty"`
	Sign        string `json:"sign,omitempty"`
}

// Данные о создании
type creation struct {
	User userLink `json:"user"`
	Date int64    `json:"date"`
}

// Статистика темы
type topicStats struct {
	MessageCount uint32 `json:"message_count"`
	ViewsCount   uint32 `json:"views_count"`
}

// Статистика сообщения
type messageStats struct {
	PlusCount  uint16 `json:"plus_count"`
	MinusCount uint16 `json:"minus_count"`
}

func getForumBlocks(dbForums []DbForum, dbModerators map[uint16][]DbModerator) []ForumBlock {
	var forumBlocks []ForumBlock

	currentForumBlockId := uint16(0) // f_forum_block.id начинаются с 1
	for _, dbForum := range dbForums {
		if dbForum.ForumBlockId != currentForumBlockId {
			forumBlock := ForumBlock{
				Id:     dbForum.ForumBlockId,
				Title:  dbForum.ForumBlockName,
				Forums: []Forum{},
			}
			forumBlocks = append(forumBlocks, forumBlock)
			currentForumBlockId = dbForum.ForumBlockId
		}
	}
	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockId == forumBlocks[index].Id {
				var moderators []userLink
				for _, dbModerator := range dbModerators[dbForum.ForumId] {
					userLink := userLink{
						Id:    dbModerator.UserId,
						Login: dbModerator.Login,
					}
					moderators = append(moderators, userLink)
				}
				forum := Forum{
					Id:          dbForum.ForumId,
					Title:       dbForum.Name,
					Description: dbForum.Description,
					Moderators:  moderators,
					Stats: forumStats{
						TopicCount:   dbForum.TopicCount,
						MessageCount: dbForum.MessageCount,
					},
					LastMessage: lastMessage{
						Id: dbForum.LastMessageId,
						Topic: topicLink{
							Id:    dbForum.LastTopicId,
							Title: dbForum.LastTopicName,
						},
						User: userLink{
							Id:    dbForum.LastUserId,
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

	return forumBlocks
}

func getForumTopics(dbTopics []DbForumTopic) []ForumTopic {
	//noinspection GoPreferNilSlice
	topics := []ForumTopic{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbTopic := range dbTopics {
		topic := ForumTopic{
			Id:    dbTopic.TopicId,
			Title: dbTopic.Name,
			Type:  dbTopic.TopicTypeId,
			Creation: creation{
				User: userLink{
					Id:    dbTopic.UserId,
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
				Id: dbTopic.LastMessageId,
				User: userLink{
					Id:    dbTopic.LastUserId,
					Login: dbTopic.LastUserName,
				},
				Date: dbTopic.LastMessageDate.Unix(),
			},
		}
		topics = append(topics, topic)
	}

	return topics
}

func getTopicMessages(dbMessages []DbForumMessage) []TopicMessage {
	//noinspection GoPreferNilSlice
	messages := []TopicMessage{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbMessage := range dbMessages {
		text := dbMessage.MessageText
		if dbMessage.IsCensored {
			text = ""
		}
		message := TopicMessage{
			Id: dbMessage.MessageId,
			Creation: creation{
				User: userLink{
					Id:          dbMessage.UserId,
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

	return messages
}
