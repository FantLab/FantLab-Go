package converters

import (
	"fantlab/base/protobuf/pbutils"
	"fantlab/core/config"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"time"
)

func GetForumBlocks(dbForumsResponse db.ForumsDBResponse, userId uint64, cfg *config.AppConfig) *pb.Forum_ForumBlocksResponse {
	var forumBlocks []*pb.Forum_ForumBlock

	currentForumBlockId := uint64(0) // f_forum_block.id начинаются с 1

	for _, dbForum := range dbForumsResponse.Forums {
		if dbForum.ForumBlockId != currentForumBlockId {
			forumBlock := pb.Forum_ForumBlock{
				Id:     dbForum.ForumBlockId,
				Title:  dbForum.ForumBlockName,
				Forums: []*pb.Forum_Forum{},
			}
			forumBlocks = append(forumBlocks, &forumBlock)
			currentForumBlockId = dbForum.ForumBlockId
		}
	}

	for _, dbForum := range dbForumsResponse.Forums {
		for index := range forumBlocks {
			if dbForum.ForumBlockId == forumBlocks[index].GetId() {
				var notModeratedTopicCount uint64
				var moderators []*pb.Common_UserLink

				for _, dbModerator := range dbForumsResponse.Moderators[dbForum.ForumId] {
					gender := helpers.GetGender(dbModerator.UserId, dbModerator.Sex)
					avatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbModerator.UserId, dbModerator.PhotoNumber)

					userLink := &pb.Common_UserLink{
						Id:     dbModerator.UserId,
						Login:  dbModerator.Login,
						Gender: gender,
						Avatar: avatar,
					}
					moderators = append(moderators, userLink)

					if userId != 0 && dbModerator.UserId == userId {
						// Количество неотмодерированных тем в форуме видно только модераторам этого форума
						notModeratedTopicCount = dbForum.NotModeratedTopicCount
					}
				}

				gender := helpers.GetGender(dbForum.LastMessageUserId, dbForum.LastMessageUserSex)
				avatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbForum.LastMessageUserId, dbForum.LastMessageUserPhotoNumber)

				forum := pb.Forum_Forum{
					Id:               dbForum.ForumId,
					Title:            dbForum.Name,
					ForumDescription: dbForum.Description,
					Moderators:       moderators,
					Stats: &pb.Forum_Forum_Stats{
						TopicCount:             dbForum.TopicCount,
						NotModeratedTopicCount: notModeratedTopicCount,
						MessageCount:           dbForum.MessageCount,
						NotReadMessageCount:    dbForumsResponse.NotReadMessageCounts[dbForum.ForumId],
					},
					LastMessage: &pb.Forum_LastMessage{
						Id: dbForum.LastMessageId,
						Topic: &pb.Forum_TopicLink{
							Id:    dbForum.LastTopicId,
							Title: dbForum.LastTopicName,
						},
						User: &pb.Common_UserLink{
							Id:     dbForum.LastMessageUserId,
							Login:  dbForum.LastMessageUserLogin,
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

func GetForumTopics(dbResponse db.ForumTopicsDBResponse, page, limit uint64, cfg *config.AppConfig) *pb.Forum_ForumTopicsResponse {
	//noinspection GoPreferNilSlice
	topics := []*pb.Forum_Topic{}

	subscribedTopics := make(map[uint64]bool, len(dbResponse.SubscribedTopics))
	for _, subscribedTopicId := range dbResponse.SubscribedTopics {
		subscribedTopics[subscribedTopicId] = true
	}

	var moderators []*pb.Common_UserLink

	for _, dbModerator := range dbResponse.Moderators {
		gender := helpers.GetGender(dbModerator.UserId, dbModerator.Sex)
		avatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbModerator.UserId, dbModerator.PhotoNumber)

		userLink := &pb.Common_UserLink{
			Id:     dbModerator.UserId,
			Login:  dbModerator.Login,
			Gender: gender,
			Avatar: avatar,
		}
		moderators = append(moderators, userLink)
	}

	for _, dbTopic := range dbResponse.Topics {
		var topicType pb.Forum_Topic_Type
		if dbTopic.TopicTypeId == 2 {
			topicType = pb.Forum_Topic_POLL
		} else {
			topicType = pb.Forum_Topic_TOPIC
		}

		creationUserGender := helpers.GetGender(dbTopic.UserId, dbTopic.UserSex)
		creationUserAvatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbTopic.UserId, dbTopic.UserPhotoNumber)

		lastMessageUserGender := helpers.GetGender(dbTopic.LastMessageUserId, dbTopic.LastMessageUserSex)
		lastMessageUserAvatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbTopic.LastMessageUserId, dbTopic.LastMessageUserPhotoNumber)

		topicNotReadInfo := dbResponse.TopicsNotReadInfo[dbTopic.TopicId]

		topic := &pb.Forum_Topic{
			Id:        dbTopic.TopicId,
			Title:     dbTopic.Name,
			TopicType: topicType,
			Creation: &pb.Common_Creation{
				User: &pb.Common_UserLink{
					Id:     dbTopic.UserId,
					Login:  dbTopic.UserLogin,
					Gender: creationUserGender,
					Avatar: creationUserAvatar,
				},
				Date: pbutils.TimestampProto(dbTopic.DateOfAdd),
			},
			IsClosed:       dbTopic.IsClosed != 0,
			IsPinned:       dbTopic.IsPinned != 0,
			IsNotModerated: dbTopic.Moderated == 0,
			IsSubscribed:   subscribedTopics[dbTopic.TopicId],
			Stats: &pb.Forum_Topic_Stats{
				MessageCount:        dbTopic.MessageCount,
				NotReadMessageCount: topicNotReadInfo.NotReadMessageCount,
				ViewCount:           dbTopic.Views,
			},
			LastMessage: &pb.Forum_LastMessage{
				Id: dbTopic.LastMessageId,
				User: &pb.Common_UserLink{
					Id:     dbTopic.LastMessageUserId,
					Login:  dbTopic.LastMessageUserLogin,
					Gender: lastMessageUserGender,
					Avatar: lastMessageUserAvatar,
				},
				Text: dbTopic.LastMessageText,
				Date: pbutils.TimestampProto(dbTopic.LastMessageDate),
			},
			FirstNotReadMessageId: topicNotReadInfo.FirstNotReadMessageId,
		}

		topics = append(topics, topic)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalTopicsCount, limit)

	return &pb.Forum_ForumTopicsResponse{
		Moderators: moderators,
		Topics:     topics,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetTopic(dbResponse db.ForumTopicMessagesDBResponse, attaches map[uint64][]*pb.Common_Attachment,
	draftAttaches []*pb.Common_Attachment, page, limit uint64, cfg *config.AppConfig, user *pb.Auth_Claims_UserInfo,
	userCanPerformAdminActions, userCanEditOwnForumMessages bool) *pb.Forum_ForumTopicResponse {

	topic := &pb.Forum_Topic{
		Id:           dbResponse.Topic.TopicId,
		Title:        dbResponse.Topic.TopicName,
		IsClosed:     dbResponse.Topic.IsClosed == 1,
		IsSubscribed: dbResponse.IsUserSubscribed,
	}

	forum := &pb.Forum_Forum{
		Id:    dbResponse.Topic.ForumId,
		Title: dbResponse.Topic.ForumName,
	}

	var poll *pb.Forum_Poll

	if dbResponse.Topic.TopicTypeId == 2 /* опрос */ {
		var answerOptions []*pb.Forum_Poll_AnswerOption

		showPollResults := dbResponse.IsUserTopicAnswerExists || dbResponse.Topic.IsClosed == 1

		for _, answer := range dbResponse.TopicAnswers {
			var voterCount uint64
			if showPollResults {
				voterCount = answer.Choices
			}

			answerOptions = append(answerOptions, &pb.Forum_Poll_AnswerOption{
				Text:       answer.Name,
				VoterCount: voterCount,
			})
		}

		var voters []*pb.Common_UserLink

		if showPollResults {
			for _, user := range dbResponse.TopicAnsweredUsers {
				voters = append(voters, &pb.Common_UserLink{
					Id:    user.UserId,
					Login: user.Login,
				})
			}
		}

		poll = &pb.Forum_Poll{
			AnswerOptions: answerOptions,
			Voters:        voters,
		}
	}

	var pinnedMessage *pb.Forum_TopicMessage

	if dbResponse.Topic.IsFirstMessagePinned == 1 {
		var dbPinnedMessage db.ForumMessage
		for _, message := range dbResponse.Messages {
			if message.Number == 1 {
				dbPinnedMessage = message
				break
			}
		}
		info := dbResponse.AdditionalInfos[dbPinnedMessage.MessageId]
		pinnedMessage = convertMessage(&dbPinnedMessage, attaches, info, cfg, user, userCanPerformAdminActions,
			userCanEditOwnForumMessages)
	}

	//noinspection GoPreferNilSlice
	messages := []*pb.Forum_TopicMessage{}

	for _, dbMessage := range dbResponse.Messages {
		info := dbResponse.AdditionalInfos[dbMessage.MessageId]
		message := convertMessage(&dbMessage, attaches, info, cfg, user, userCanPerformAdminActions,
			userCanEditOwnForumMessages)
		messages = append(messages, message)
	}

	var messageDraft *pb.Forum_TopicMessageDraft
	if dbResponse.MessageDraft != (db.ForumMessageDraft{}) {
		messageDraft = convertMessageDraft(&dbResponse.MessageDraft, draftAttaches, cfg)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalMessagesCount, limit)

	return &pb.Forum_ForumTopicResponse{
		Topic:         topic,
		Forum:         forum,
		Poll:          poll,
		PinnedMessage: pinnedMessage,
		Messages:      messages,
		MessageDraft:  messageDraft,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetForumTopicMessage(dbMessage *db.ForumMessage, attaches []*pb.Common_Attachment,
	info db.ForumAdditionalMessageInfo, cfg *config.AppConfig, user *pb.Auth_Claims_UserInfo,
	userCanPerformAdminActions, userCanEditOwnForumMessages bool) *pb.Forum_ForumMessageResponse {

	attachesMap := map[uint64][]*pb.Common_Attachment{}
	attachesMap[dbMessage.MessageId] = attaches

	message := convertMessage(dbMessage, attachesMap, info, cfg, user, userCanPerformAdminActions,
		userCanEditOwnForumMessages)

	return &pb.Forum_ForumMessageResponse{
		Message: message,
	}
}

func convertMessage(dbMessage *db.ForumMessage, attaches map[uint64][]*pb.Common_Attachment,
	info db.ForumAdditionalMessageInfo, cfg *config.AppConfig, user *pb.Auth_Claims_UserInfo,
	userCanPerformAdminActions, userCanEditOwnForumMessages bool) *pb.Forum_TopicMessage {

	text := dbMessage.MessageText

	userLoggedIn := user != nil
	userIsForumModerator := userLoggedIn && info.ForumModerators[user.UserId]

	isRed := dbMessage.IsRed != 0
	isCensored := dbMessage.IsCensored != 0

	if isCensored {
		text = cfg.CensorshipText
	} else if !(dbMessage.IsUserApproved == 1 || (userLoggedIn && dbMessage.UserId == user.UserId) ||
		userIsForumModerator || userCanPerformAdminActions) {
		// Если автор сообщения не подтвержден, текущий пользователь - не этот автор, не модератор и не имеет доступа к
		// админским функциям, текст сообщения заменяем на заглушку о премодерации сообщения
		text = cfg.PreModerationText
	} else {
		if userLoggedIn {
			if user.SmilesDisabled {
				// TODO Если смайлы отключены, их надо не вырезать, а заменять на алиасы (см. main.cfg#all_smiles_txt)
				text = cfg.Smiles.RemoveFromString(text)
			}
			if user.ImagesDisabled {
				text = helpers.ReplaceImagesWithUrls(text, cfg.ImageReplacementLinkText)
			}
		}
	}

	isUnread := userLoggedIn && dbMessage.DateOfAdd.After(info.DateOfTopicRead)

	var attachments []*pb.Common_Attachment
	if userLoggedIn && !isCensored {
		// Аттачи показываем только залогиненным и только у нецензурированных сообщений
		attachments = attaches[dbMessage.MessageId]
	}

	gender := helpers.GetGender(dbMessage.UserId, dbMessage.UserSex)
	avatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbMessage.UserId, dbMessage.UserPhotoNumber)

	// Пояснения по логике можно найти в коде endpoint-ов edit_forum_message, delete_forum_message
	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > cfg.MaxForumMessageEditTimeout
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages
	canUserEditFirstTopicMessage := dbMessage.Number == 1 && info.TopicStarterCanEditFirstMessage
	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	canEdit := userLoggedIn &&
		((user.UserId == dbMessage.UserId && (canUserEditMessage || canUserEditFirstTopicMessage) && isMessageEditable) ||
			userIsForumModerator || info.OnlyForAdminsForum)

	canDelete := userLoggedIn &&
		((user.UserId == dbMessage.UserId && canUserEditMessage && isMessageEditable) ||
			(userIsForumModerator && !info.ForumModerators[dbMessage.UserId]))

	var rating int64
	var canVoteMinus bool
	var canVotePlus bool

	var isForumWithEnabledRating bool
	for _, forumId := range cfg.ForumsWithEnabledRating {
		if dbMessage.ForumId == forumId {
			isForumWithEnabledRating = true
			break
		}
	}

	if isForumWithEnabledRating {
		var isForumWithDisabledMinuses bool
		for _, forumId := range cfg.ForumsWithDisabledMinuses {
			if dbMessage.ForumId == forumId {
				isForumWithDisabledMinuses = true
				break
			}
		}

		if ((userLoggedIn && !user.ForumMessagesRatingDisabled) || !userLoggedIn) && !isRed {
			if isForumWithDisabledMinuses {
				rating = int64(dbMessage.VotePlus)
			} else {
				rating = int64(dbMessage.VotePlus - dbMessage.VoteMinus)
			}
		}

		if userLoggedIn && !user.ForumMessagesRatingDisabled {
			var isReadOnlyUser bool
			for _, userId := range cfg.ReadOnlyForumUsers[dbMessage.ForumId] {
				if user.UserId == userId {
					isReadOnlyUser = true
					break
				}
			}

			canVote := !isCensored && !isRed && !info.IsVotedByUser && !isReadOnlyUser &&
				!(dbMessage.UserId == user.UserId || user.Class < pb.Common_USERCLASS_AUTHORITY)

			canVoteMinus = canVote && !isForumWithDisabledMinuses
			canVotePlus = canVote
		}
	}

	canDeleteVotes := userLoggedIn && userIsForumModerator && isForumWithEnabledRating

	return &pb.Forum_TopicMessage{
		Id: dbMessage.MessageId,
		Creation: &pb.Common_Creation{
			User: &pb.Common_UserLink{
				Id:     dbMessage.UserId,
				Login:  dbMessage.UserLogin,
				Gender: gender,
				Avatar: avatar,
				Class:  helpers.UserClassMap[dbMessage.UserClass],
				Sign:   dbMessage.UserSign,
			},
			Date: pbutils.TimestampProto(dbMessage.DateOfAdd),
		},
		Text:        text,
		IsCensored:  isCensored,
		IsUnread:    isUnread,
		Attachments: attachments,
		Stats: &pb.Forum_TopicMessage_Stats{
			Rating: rating,
		},
		Rights: &pb.Forum_TopicMessage_Rights{
			CanEdit:        canEdit,
			CanDelete:      canDelete,
			CanVoteMinus:   canVoteMinus,
			CanVotePlus:    canVotePlus,
			CanDeleteVotes: canDeleteVotes,
		},
	}
}

func GetForumTopicMessageDraft(dbMessageDraft *db.ForumMessageDraft, attaches []*pb.Common_Attachment,
	cfg *config.AppConfig) *pb.Forum_ForumMessageDraftResponse {

	messageDraft := convertMessageDraft(dbMessageDraft, attaches, cfg)

	return &pb.Forum_ForumMessageDraftResponse{
		MessageDraft: messageDraft,
	}
}

func convertMessageDraft(dbMessageDraft *db.ForumMessageDraft, attaches []*pb.Common_Attachment,
	cfg *config.AppConfig) *pb.Forum_TopicMessageDraft {

	gender := helpers.GetGender(dbMessageDraft.UserID, dbMessageDraft.Sex)
	avatar := helpers.GetUserAvatarUrl(cfg.BaseImageUrl, dbMessageDraft.UserID, dbMessageDraft.PhotoNumber)

	return &pb.Forum_TopicMessageDraft{
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
		Text:        dbMessageDraft.Message,
		Attachments: attaches,
	}
}
