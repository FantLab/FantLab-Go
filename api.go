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
	Id              uint16     `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Moderators      []UserLink `json:"moderators"`
	TopicCount      uint32     `json:"topic_count"`
	MessageCount    uint32     `json:"message_count"`
	LastTopic       TopicLink  `json:"last_topic"`
	LastMessageUser UserLink   `json:"last_message_user"`
	LastMessageDate time.Time  `json:"last_message_date"`
}

// Ссылка на тему форума
type TopicLink struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
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
			forumBlocks = append(forumBlocks, ForumBlock{
				Id:     dbForum.ForumBlockId,
				Name:   dbForum.ForumBlockName,
				Forums: []Forum{},
			})
			currentForumBlockId = dbForum.ForumBlockId
		}
	}
	for _, dbForum := range dbForums {
		for index := range forumBlocks {
			if dbForum.ForumBlockId == forumBlocks[index].Id {
				var moderators []UserLink
				for _, dbModerator := range dbModerators {
					if dbModerator.ForumId == dbForum.ForumId {
						moderators = append(moderators, UserLink{
							Id:    dbModerator.UserId,
							Login: dbModerator.Login,
						})
					}
				}
				forumBlocks[index].Forums = append(forumBlocks[index].Forums, Forum{
					Id:           dbForum.ForumId,
					Name:         dbForum.Name,
					Description:  dbForum.Description,
					Moderators:   moderators,
					TopicCount:   dbForum.TopicCount,
					MessageCount: dbForum.MessageCount,
					LastTopic: TopicLink{
						Id:   dbForum.LastTopicId,
						Name: dbForum.LastTopicName,
					},
					LastMessageUser: UserLink{
						Id:    dbForum.LastUserId,
						Login: dbForum.LastUserName,
					},
					LastMessageDate: dbForum.LastMessageDate,
				})
				break
			}
		}
	}
	return forumBlocks
}
