package main

// Блок форумов
type ForumBlock struct {
	Id     uint16  `json:"-"`
	Title  string  `json:"block_title"`
	Forums []Forum `json:"forums"`
}

// Форум
type Forum struct {
	Id          uint16           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Moderators  []userLink       `json:"moderators"`
	Stats       forumStats       `json:"stats"`
	LastMessage lastForumMessage `json:"last_message"`
}

// Тема
type Topic struct {
	Id          uint32           `json:"id"`
	Title       string           `json:"title"`
	Type        uint16           `json:"type"`
	Creation    topicCreation    `json:"creation"`
	IsClosed    bool             `json:"is_closed"`
	IsPinned    bool             `json:"is_pinned"`
	Stats       topicStats       `json:"stats"`
	LastMessage lastTopicMessage `json:"last_message"`
}

// Статистика форума
type forumStats struct {
	TopicCount   uint32 `json:"topic_count"`
	MessageCount uint32 `json:"message_count"`
}

// Последнее сообщение в форуме
type lastForumMessage struct {
	Id    uint32    `json:"id"`
	Topic topicLink `json:"topic"`
	User  userLink  `json:"user"`
	Date  int64     `json:"date"`
}

// Ссылка на тему форума
type topicLink struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

// Ссылка на пользователя
type userLink struct {
	Id    uint32 `json:"id"`
	Login string `json:"login"`
}

// Данные о создании темы
type topicCreation struct {
	User userLink `json:"user"`
	Date int64    `json:"date"`
}

// Статистика темы
type topicStats struct {
	MessageCount uint32 `json:"message_count"`
	ViewsCount   uint32 `json:"views_count"`
}

// Последнее сообщение в теме
type lastTopicMessage struct {
	Id   uint32   `json:"id"`
	User userLink `json:"user"`
	Date int64    `json:"date"`
}

func getForumBlocks(dbForums []DbForum, dbModerators []DbModerator) []ForumBlock {
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
				for _, dbModerator := range dbModerators {
					if dbModerator.ForumId == dbForum.ForumId {
						userLink := userLink{
							Id:    dbModerator.UserId,
							Login: dbModerator.Login,
						}
						moderators = append(moderators, userLink)
					}
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
					LastMessage: lastForumMessage{
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

func getTopics(dbTopics []DbTopic) []Topic {
	//noinspection GoPreferNilSlice
	topics := []Topic{} // возвращаем в случае отсутствия результатов пустой массив

	for _, dbTopic := range dbTopics {
		topics = append(topics, Topic{
			Id:    dbTopic.TopicId,
			Title: dbTopic.Name,
			Type:  dbTopic.TopicTypeId,
			Creation: topicCreation{
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
			LastMessage: lastTopicMessage{
				Id: dbTopic.LastMessageId,
				User: userLink{
					Id:    dbTopic.LastUserId,
					Login: dbTopic.LastUserName,
				},
				Date: dbTopic.LastMessageDate.Unix(),
			},
		})
	}

	return topics
}
