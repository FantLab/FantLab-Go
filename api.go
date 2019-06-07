package main

import (
	"time"
)

// Блок форумов
type ForumBlock struct {
	Id     uint16  `json:"-"`
	Name   string  `json:"block_name"`
	Forums []Forum `json:"forums"`
}

// Форум
type Forum struct {
	Id          uint16           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Moderators  []UserLink       `json:"moderators"`
	Stats       ForumStats       `json:"stats"`
	LastMessage LastForumMessage `json:"last_message"`
}

// Статистика форума
type ForumStats struct {
	TopicCount   uint32 `json:"topic_count"`
	MessageCount uint32 `json:"message_count"`
}

// Последнее сообщение в форуме
type LastForumMessage struct {
	Topic TopicLink `json:"topic"`
	User  UserLink  `json:"user"`
	Date  time.Time `json:"date"`
}

// Ссылка на тему форума
type TopicLink struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

// Ссылка на пользователя
type UserLink struct {
	Id    uint32 `json:"id"`
	Login string `json:"login"`
}

func getForumBlocks(dbForums []DbForum, dbModerators []DbModerator) []ForumBlock {
	// todo better algorithm
	var forumBlocks []ForumBlock
	currentForumBlockId := uint16(0)
	for _, dbForum := range dbForums {
		if dbForum.ForumBlockId != currentForumBlockId {
			forumBlock := ForumBlock{
				Id:     dbForum.ForumBlockId,
				Name:   dbForum.ForumBlockName,
				Forums: []Forum{},
			}
			forumBlocks = append(forumBlocks, forumBlock)
			currentForumBlockId = dbForum.ForumBlockId
		}
	}
	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockId == forumBlocks[index].Id {
				var moderators []UserLink
				for _, dbModerator := range dbModerators {
					if dbModerator.ForumId == dbForum.ForumId {
						userLink := UserLink{
							Id:    dbModerator.UserId,
							Login: dbModerator.Login,
						}
						moderators = append(moderators, userLink)
					}
				}
				forum := Forum{
					Id:          dbForum.ForumId,
					Name:        dbForum.Name,
					Description: dbForum.Description,
					Moderators:  moderators,
					Stats: ForumStats{
						TopicCount:   dbForum.TopicCount,
						MessageCount: dbForum.MessageCount,
					},
					LastMessage: LastForumMessage{
						Topic: TopicLink{
							Id:    dbForum.LastTopicId,
							Title: dbForum.LastTopicName,
						},
						User: UserLink{
							Id:    dbForum.LastUserId,
							Login: dbForum.LastUserName,
						},
						Date: dbForum.LastMessageDate,
					},
				}
				forumBlocks[index].Forums = append(forumBlocks[index].Forums, forum)
				break
			}
		}
	}
	return forumBlocks
}
