package db

import (
	"context"
	"fantlab/core/db/queries"
	"fantlab/core/helpers"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqlbuilder"
)

type ForumInList struct {
	ForumId                    uint64    `db:"forum_id"`
	Name                       string    `db:"name"`
	Description                string    `db:"description"`
	TopicCount                 uint64    `db:"topic_count"`
	MessageCount               uint64    `db:"message_count"`
	LastTopicId                uint64    `db:"last_topic_id"`
	LastTopicName              string    `db:"last_topic_name"`
	LastMessageId              uint64    `db:"last_message_id"`
	LastMessageUserId          uint64    `db:"last_message_user_id"`
	LastMessageUserLogin       string    `db:"last_message_user_login"`
	LastMessageUserSex         uint8     `db:"last_message_user_sex"`
	LastMessageUserPhotoNumber uint64    `db:"last_message_user_photo_number"`
	LastMessageText            string    `db:"last_message_text"`
	LastMessageDate            time.Time `db:"last_message_date"`
	ForumBlockId               uint64    `db:"forum_block_id"`
	ForumBlockName             string    `db:"forum_block_name"`
	NotModeratedTopicCount     uint64    `db:"not_moderated_topic_count"`
}

type Forum struct {
	ForumId uint64 `db:"forum_id"`
	Name    string `db:"name"`
	// Доступ в форум только админам. NOTE На самом деле это поле крайне опасно. Если не делать явную проверку на наличие
	// доступа к такому форуму, у любого пользователя появляется возможность редактировать сообщения в нем.
	// В Perl-бэке такая уязвимость была: https://github.com/parserpro/fantlab/issues/954,
	// https://github.com/parserpro/fantlab/issues/952
	OnlyForAdmins uint8 `db:"only_for_admins"`
	ForumClosed   uint8 `db:"forum_closed"`
}

type ForumTopic struct {
	TopicId                    uint64    `db:"topic_id"`
	ForumId                    uint64    `db:"forum_id"`
	Name                       string    `db:"name"`
	DateOfAdd                  time.Time `db:"date_of_add"`
	Views                      uint64    `db:"views"`
	UserId                     uint64    `db:"user_id"`
	UserLogin                  string    `db:"user_login"`
	UserSex                    uint8     `db:"user_sex"`
	UserPhotoNumber            uint64    `db:"user_photo_number"`
	TopicTypeId                uint64    `db:"topic_type_id"`
	IsClosed                   uint8     `db:"is_closed"`
	IsPinned                   uint8     `db:"is_pinned"`
	MessageCount               uint64    `db:"message_count"`
	LastMessageId              uint64    `db:"last_message_id"`
	LastMessageUserId          uint64    `db:"last_message_user_id"`
	LastMessageUserLogin       string    `db:"last_message_user_login"`
	LastMessageUserSex         uint8     `db:"last_message_user_sex"`
	LastMessageUserPhotoNumber uint64    `db:"last_message_user_photo_number"`
	LastMessageText            string    `db:"last_message_text"`
	LastMessageDate            time.Time `db:"last_message_date"`
	Moderated                  uint8     `db:"moderated"`
}

type ShortForumTopic struct {
	TopicId              uint64 `db:"topic_id"`
	TopicName            string `db:"topic_name"`
	IsFirstMessagePinned uint8  `db:"is_first_message_pinned"`
	TopicTypeId          uint8  `db:"topic_type_id"`
	IsClosed             uint8  `db:"is_closed"`
	IsEditTopicStarter   uint8  `db:"is_edit_topicstarter"`
	ForumId              uint64 `db:"forum_id"`
	ForumName            string `db:"forum_name"`
}

type ForumMessage struct {
	MessageId       uint64    `db:"message_id"`
	TopicId         uint64    `db:"topic_id"`
	ForumId         uint64    `db:"forum_id"`
	DateOfAdd       time.Time `db:"date_of_add"`
	IsRed           uint8     `db:"is_red"` // модераторское?
	IsCensored      uint8     `db:"is_censored"`
	VotePlus        uint64    `db:"vote_plus"`
	VoteMinus       uint64    `db:"vote_minus"`
	Number          uint64    `db:"number"`
	MessageText     string    `db:"message_text"`
	UserId          uint64    `db:"user_id"`
	UserLogin       string    `db:"login"`
	UserSex         uint8     `db:"sex"`
	UserPhotoNumber uint64    `db:"photo_number"`
	UserClass       uint8     `db:"user_class"`
	UserSign        string    `db:"sign"`
	IsUserApproved  uint8     `db:"approved"`
}

type ForumMessageAttachment struct {
	MessageId uint64 `db:"message_id"`
	FileName  string `db:"file_name"`
	FileSize  uint64 `db:"file_size"`
}

// Вариант ответа в опросе
type ForumTopicAnswer struct {
	Number  uint64 `db:"number"`
	Name    string `db:"name"`
	Choices uint64 `db:"choices"`
}

type ForumModerator struct {
	ForumId     uint64 `db:"forum_id"`
	UserId      uint64 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint64 `db:"photo_number"`
}

type ForumTopicNotReadInfo struct {
	TopicId               uint64 `db:"topic_id"`
	FirstNotReadMessageId uint64 `db:"first_not_read_message_id"`
	NotReadMessageCount   uint64 `db:"not_read_message_count"`
}

type ForumAdditionalMessageInfo struct {
	MessageId                       uint64
	IsVotedByUser                   bool
	IsWarned                        bool
	IsModerCalled                   bool
	TopicStarterCanEditFirstMessage bool
	OnlyForAdminsForum              bool
	DateOfTopicRead                 time.Time
	ForumModerators                 map[uint64]bool
}

type ForumsDBResponse struct {
	Forums               []ForumInList
	Moderators           map[uint64][]ForumModerator
	NotReadMessageCounts map[uint64]uint64
}

type ForumTopicsDBResponse struct {
	Moderators        []ForumModerator
	Topics            []ForumTopic
	SubscribedTopics  []uint64
	TopicsNotReadInfo map[uint64]ForumTopicNotReadInfo
	TotalTopicsCount  uint64
}

type ForumTopicMessagesDBResponse struct {
	Topic                   ShortForumTopic
	Messages                []ForumMessage
	Attachments             []ForumMessageAttachment
	AdditionalInfos         map[uint64]ForumAdditionalMessageInfo
	TotalMessagesCount      uint64
	MessageDraft            ForumMessageDraft
	IsUserSubscribed        bool
	IsUserTopicAnswerExists bool
	TopicAnswers            []ForumTopicAnswer
	TopicAnsweredUsers      []User
}

func (db *DB) FetchForums(ctx context.Context, userId uint64, availableForums []uint64) (ForumsDBResponse, error) {
	var forums []ForumInList
	var moderators []ForumModerator
	var notModeratedTopicIds []uint64
	notReadMessageCountMap := map[uint64]uint64{}

	err := codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForums).WithArgs(availableForums).FlatArgs(), &forums)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetModerators), &moderators)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetNotModeratedTopicIds), &notModeratedTopicIds)
		},
		func() error {
			if len(notModeratedTopicIds) == 0 {
				notModeratedTopicIds = append(notModeratedTopicIds, 0)
			}
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetNotReadMessageCounts).WithArgs(userId, notModeratedTopicIds).FlatArgs(), &notReadMessageCountMap)
		},
	)

	if err != nil {
		return ForumsDBResponse{}, err
	}

	moderatorsMap := map[uint64][]ForumModerator{}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumId] = append(moderatorsMap[moderator.ForumId], moderator)
	}

	response := ForumsDBResponse{
		Forums:               forums,
		Moderators:           moderatorsMap,
		NotReadMessageCounts: notReadMessageCountMap,
	}

	return response, nil
}

func (db *DB) FetchForumExists(ctx context.Context, forumId uint64, availableForums []uint64) (bool, error) {
	var isExists uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumExists).WithArgs(forumId, availableForums).FlatArgs(), &isExists)

	if err != nil {
		return false, err
	}

	return isExists == 1, nil
}

func (db *DB) FetchForum(ctx context.Context, forumId uint64) (Forum, error) {
	var forum Forum

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForum).WithArgs(forumId), &forum)

	if err != nil {
		return Forum{}, err
	}

	return forum, nil
}

func (db *DB) FetchForumTopics(ctx context.Context, userId, forumId, limit, offset uint64) (ForumTopicsDBResponse, error) {
	var moderators []ForumModerator
	var topics []ForumTopic
	var topicCount uint64
	var subscribedTopics []uint64

	moderatedState := []uint8{1}
	var topicIds []uint64

	type topicDateOfRead struct {
		TopicId    uint64    `db:"topic_id"`
		DateOfRead time.Time `db:"date_of_read"`
	}

	var topicsDatesOfRead []topicDateOfRead

	topicsNotReadInfo := map[uint64]ForumTopicNotReadInfo{}

	err := codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumModerators).WithArgs(forumId), &moderators)
		},
		func() error {
			for _, moderator := range moderators {
				if moderator.UserId == userId {
					// Если текущий пользователь - модератор, он должен видеть и неотмодерированные темы
					moderatedState = append(moderatedState, 0)
					break
				}
			}
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumTopics).WithArgs(forumId, moderatedState, userId, limit, offset).FlatArgs(), &topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumTopicCount).WithArgs(forumId, moderatedState, userId).FlatArgs(), &topicCount)
		},
		func() error {
			for _, topic := range topics {
				topicIds = append(topicIds, topic.TopicId)
			}
			if userId != 0 && len(topicIds) != 0 {
				return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicsSubscriptions).WithArgs(userId, topicIds).FlatArgs(), &subscribedTopics)
			} else {
				return nil
			}
		},
		func() error {
			if userId != 0 && len(topicIds) != 0 {
				return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicsDatesOfRead).WithArgs(userId, topicIds).FlatArgs(), &topicsDatesOfRead)
			} else {
				return nil
			}
		},
		func() error {
			if userId != 0 {
				var topicNotReadInfo ForumTopicNotReadInfo
				for _, topicDateOfRead := range topicsDatesOfRead {
					err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetNotReadTopicInfo).WithArgs(topicDateOfRead.TopicId, topicDateOfRead.DateOfRead), &topicNotReadInfo)
					if err2 != nil && !IsNotFoundError(err2) {
						return err2
					}
					topicsNotReadInfo[topicNotReadInfo.TopicId] = topicNotReadInfo
				}
			}
			return nil
		},
	)

	if err != nil {
		return ForumTopicsDBResponse{}, err
	}

	response := ForumTopicsDBResponse{
		Moderators:        moderators,
		Topics:            topics,
		SubscribedTopics:  subscribedTopics,
		TopicsNotReadInfo: topicsNotReadInfo,
		TotalTopicsCount:  topicCount,
	}

	return response, nil
}

func (db *DB) FetchForumTopicExists(ctx context.Context, topicId uint64, availableForums []uint64) (bool, error) {
	var isExists uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicExists).WithArgs(topicId, availableForums).FlatArgs(), &isExists)

	if err != nil {
		return false, err
	}

	return isExists == 1, nil
}

func (db *DB) FetchForumTopic(ctx context.Context, topicId uint64) (*ForumTopic, error) {
	var topic ForumTopic

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopic).WithArgs(topicId), &topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchForumTopicShort(ctx context.Context, topicId uint64) (ShortForumTopic, error) {
	var shortTopic ShortForumTopic

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicShort).WithArgs(topicId), &shortTopic)

	if err != nil {
		return ShortForumTopic{}, err
	}

	return shortTopic, nil
}

func (db *DB) FetchAdditionalMessageInfo(ctx context.Context, messageId, topicId, forumId, userId uint64) (ForumAdditionalMessageInfo, error) {
	var forumModerators []ForumModerator
	var votedMessageIds []uint64
	var warnedMessageIds []uint64
	var moderCalledMessageIds []uint64
	var forum Forum
	var shortTopic ShortForumTopic

	type dateOfRead struct {
		DateOfRead time.Time `db:"date_of_read"`
	}
	var dateOfTopicRead dateOfRead

	err := codeflow.Try(
		func() error { // Получаем список модераторов форума
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumModerators).WithArgs(forumId), &forumModerators)
		},
		func() error { // Получаем список сообщений, которым выписаны предупреждения
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetWarnedMessageIds).WithArgs(messageId), &warnedMessageIds)
		},
		func() error { // Получаем список сообщений, которым вызван модератор
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetModerCalledMessageIds).WithArgs(messageId), &moderCalledMessageIds)
		},
		func() error {
			// Получаем данные о форуме (по идее, к этому моменту мы уже проверили его на существование, так что
			// IsNotFoundError не может быть выкинут)
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForum).WithArgs(forumId), &forum)
		},
		func() error { // Получаем данные о теме
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicShort).WithArgs(topicId), &shortTopic)
		},
		func() error {
			if userId != 0 {
				return codeflow.Try(
					func() error { // Получаем дату прочтения, если уже заходили в тему
						err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicReadDate).WithArgs(topicId, userId), &dateOfTopicRead)
						if IsNotFoundError(err2) {
							return nil
						}
						return err2
					},
					func() error { // Получаем список сообщений, за которые голосовал (+/-) текущий пользователь
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetVotedMessageIds).WithArgs(messageId, userId), &votedMessageIds)
					},
				)
			}
			return nil
		},
	)

	if err != nil {
		return ForumAdditionalMessageInfo{}, err
	}

	var isVotedByUser bool
	if len(votedMessageIds) > 0 {
		isVotedByUser = true
	}

	var isWarned bool
	if len(warnedMessageIds) > 0 {
		isWarned = true
	}

	var isModerCalled bool
	if len(moderCalledMessageIds) > 0 {
		isModerCalled = true
	}

	moderators := map[uint64]bool{}
	for _, forumModerator := range forumModerators {
		moderators[forumModerator.UserId] = true
	}

	info := ForumAdditionalMessageInfo{
		MessageId:                       messageId,
		IsVotedByUser:                   isVotedByUser,
		IsWarned:                        isWarned,
		IsModerCalled:                   isModerCalled,
		TopicStarterCanEditFirstMessage: shortTopic.IsEditTopicStarter == 1,
		OnlyForAdminsForum:              forum.OnlyForAdmins == 1,
		DateOfTopicRead:                 dateOfTopicRead.DateOfRead,
		ForumModerators:                 moderators,
	}

	return info, nil
}

func (db *DB) FetchTopicMessages(ctx context.Context, topicId, limit, offset uint64, sortAsc bool, userId uint64) (ForumTopicMessagesDBResponse, error) {
	var messageCount uint64
	var messageIds []uint64
	var pinnedFirstMessageId uint64
	var messages []ForumMessage
	var attachments []ForumMessageAttachment
	var shortTopic ShortForumTopic
	var forum Forum
	var forumModerators []ForumModerator
	var votedMessageIds []uint64
	var warnedMessageIds []uint64
	var moderCalledMessageIds []uint64
	var messageDraft ForumMessageDraft
	var isUserSubscribed uint8
	var isUserTopicAnswerExists uint8
	var topicAnswers []ForumTopicAnswer
	var topicAnsweredUsers []User

	type dateOfRead struct {
		DateOfRead time.Time `db:"date_of_read"`
	}
	var dateOfTopicRead dateOfRead

	err := codeflow.Try(
		func() error { // Получаем количество всех сообщений в теме
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumTopicMessageCount).WithArgs(topicId), &messageCount)
		},
		func() error { // Получаем список всех сообщений в теме
			iOffset := int64(offset)
			iLimit := int64(limit)
			iMessageCount := int64(messageCount)

			var minNumber int64
			var maxNumber int64
			if sortAsc {
				maxNumber = iOffset + iLimit
				minNumber = maxNumber - (iLimit - 1)
				if minNumber > iMessageCount {
					minNumber = iMessageCount + 1
				}
				if maxNumber > iMessageCount {
					maxNumber = iMessageCount
				}
			} else {
				maxNumber = iMessageCount - iOffset
				minNumber = maxNumber - (iLimit - 1)
				if minNumber < 0 {
					minNumber = 0
				}
				if maxNumber < 0 {
					maxNumber = -1
				}
			}

			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessageIds).WithArgs(topicId, minNumber, maxNumber), &messageIds)
		},
		func() error { // Получаем данные о теме
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicShort).WithArgs(topicId), &shortTopic)
		},
		func() error { // Получаем id запиненного сообщения, если есть
			if shortTopic.IsFirstMessagePinned == 1 {
				err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicFirstMessageId).WithArgs(topicId), &pinnedFirstMessageId)
				if err2 == nil {
					messageIds = append(messageIds, pinnedFirstMessageId)
				}
				return err2
			}
			return nil
		},
		func() error {
			if len(messageIds) > 0 {
				return codeflow.Try(
					func() error { // Получаем сами сообщения
						var sortDirection string
						if sortAsc {
							sortDirection = "ASC"
						} else {
							sortDirection = "DESC"
						}

						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessages).Inject(sortDirection).WithArgs(messageIds).FlatArgs(), &messages)
					},
					func() error { // Получаем список аттачей ко всем сообщениям темы
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessagesAttachments).WithArgs(messageIds).FlatArgs(), &attachments)
					},
					func() error { // Получаем список сообщений, которым выписаны предупреждения
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetWarnedMessageIds).WithArgs(messageIds).FlatArgs(), &warnedMessageIds)
					},
					func() error { // Получаем список сообщений, которым вызван модератор
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetModerCalledMessageIds).WithArgs(messageIds).FlatArgs(), &moderCalledMessageIds)
					},
				)
			}
			return nil
		},
		func() error { // Получаем данные о форуме
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForum).WithArgs(shortTopic.ForumId), &forum)
		},
		func() error { // Получаем список модераторов форума
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumModerators).WithArgs(shortTopic.ForumId), &forumModerators)
		},
		func() error {
			if userId != 0 {
				return codeflow.Try(
					func() error { // Получаем дату прочтения, если уже заходили в тему
						err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicReadDate).WithArgs(topicId, userId), &dateOfTopicRead)
						if IsNotFoundError(err2) {
							return nil
						}
						return err2
					},
					func() error { // Получаем черновик, если есть
						err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId), &messageDraft)
						if IsNotFoundError(err2) {
							return nil
						}
						return err2
					},
					func() error { // Получаем список сообщений, за которые голосовал (+/-) текущий пользователь
						if len(messageIds) > 0 {
							return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetVotedMessageIds).WithArgs(messageIds, userId).FlatArgs(), &votedMessageIds)
						}
						return nil
					},
					func() error { // Получаем статус подписки на тему
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicSubscriptionExists).WithArgs(topicId, userId), &isUserSubscribed)
					},
				)
			}
			return nil
		},
		func() error { // Получаем данные по опросу
			if shortTopic.TopicTypeId == 2 {
				return codeflow.Try(
					func() error { // Получаем варианты ответов
						return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicAnswers).WithArgs(topicId), &topicAnswers)
					},
					func() error {
						if userId != 0 {
							return codeflow.Try(
								func() error { // Пользователь дал ответ в опросе?
									return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetUserTopicAnswerExists).WithArgs(topicId, userId), &isUserTopicAnswerExists)
								},
								func() error {
									// Если дал или тема закрыта, получаем список проголосовавших (здесь просачивается
									// бизнес-логика, но зато одним запросом меньше)
									if isUserTopicAnswerExists == 1 || shortTopic.IsClosed == 1 {
										return db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicAnsweredUsers).WithArgs(topicId), &topicAnsweredUsers)
									}
									return nil
								},
							)
						}
						return nil
					},
				)
			}
			return nil
		},
		func() error {
			return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
				return codeflow.Try(
					func() error { // Инкрементим количество просмотров темы
						return rw.Write(ctx, sqlapi.NewQuery(queries.ForumIncrementTopicViewCount).WithArgs(topicId)).Error
					},
					func() error {
						if userId != 0 {
							var mostRecentMessageDateOfAdd time.Time
							var previousReadMessageCount uint64

							return codeflow.Try(
								func() error { // Удаляем прочитанные сообщения из списка непрочитанных
									if len(messageIds) > 0 {
										if sortAsc {
											mostRecentMessageDateOfAdd = messages[len(messages)-1].DateOfAdd
										} else {
											mostRecentMessageDateOfAdd = messages[0].DateOfAdd
										}
										return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteUserForumNewMessages).
											WithArgs(userId, topicId, messageIds, mostRecentMessageDateOfAdd).FlatArgs()).Error
									}
									return nil
								},
								func() error { // Обновляем количество новых сообщений в форуме
									return rw.Write(ctx, sqlapi.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(userId, userId)).Error
								},
								func() error { // Получаем количество прочитанных сообщений в теме
									err2 := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetUserTopicReadCount).WithArgs(userId, topicId), &previousReadMessageCount)
									if IsNotFoundError(err2) {
										return nil
									}
									return err2
								},
								func() error { // Выставляем новую дату прочтения темы
									isLastTopicPage := (sortAsc && (offset+limit >= messageCount)) ||
										(!sortAsc && (offset == 0))
									if isLastTopicPage {
										// Логика Perl-бэка
										now := time.Now()
										return rw.Write(ctx, sqlapi.NewQuery(queries.ForumInsertUserTopicReadDate).
											WithArgs(userId, topicId, now, messageCount, forum.ForumId,
												now, messageCount, forum.ForumId)).Error
									} else {
										var existsUnreadMessages bool
										if dateOfTopicRead.DateOfRead.IsZero() {
											existsUnreadMessages = true
										} else {
											for _, message := range messages {
												if message.DateOfAdd.After(dateOfTopicRead.DateOfRead) {
													existsUnreadMessages = true
													break
												}
											}
										}
										if existsUnreadMessages {
											var readMessageCount uint64
											if sortAsc {
												readMessageCount = offset + limit
											} else {
												readMessageCount = messageCount - offset
											}
											if previousReadMessageCount > readMessageCount {
												readMessageCount = previousReadMessageCount
											}
											return rw.Write(ctx, sqlapi.NewQuery(queries.ForumInsertUserTopicReadDate).
												WithArgs(userId, topicId, mostRecentMessageDateOfAdd, readMessageCount, forum.ForumId,
													mostRecentMessageDateOfAdd, readMessageCount, forum.ForumId)).Error
										}
									}
									return nil
								},
							)
						}
						return nil
					},
				)
			})
		},
	)

	if err != nil {
		return ForumTopicMessagesDBResponse{}, err
	}

	votedMessagesMap := map[uint64]bool{}
	for _, messageId := range votedMessageIds {
		votedMessagesMap[messageId] = true
	}

	warnedMessagesMap := map[uint64]bool{}
	for _, messageId := range warnedMessageIds {
		warnedMessagesMap[messageId] = true
	}

	moderCalledMessagesMap := map[uint64]bool{}
	for _, messageId := range moderCalledMessageIds {
		moderCalledMessagesMap[messageId] = true
	}

	moderators := map[uint64]bool{}
	for _, forumModerator := range forumModerators {
		moderators[forumModerator.UserId] = true
	}

	additionalInfos := map[uint64]ForumAdditionalMessageInfo{}
	for _, messageId := range messageIds {
		additionalInfos[messageId] = ForumAdditionalMessageInfo{
			MessageId:                       messageId,
			IsVotedByUser:                   votedMessagesMap[messageId],
			IsWarned:                        warnedMessagesMap[messageId],
			IsModerCalled:                   moderCalledMessagesMap[messageId],
			TopicStarterCanEditFirstMessage: shortTopic.IsEditTopicStarter == 1,
			OnlyForAdminsForum:              forum.OnlyForAdmins == 1,
			DateOfTopicRead:                 dateOfTopicRead.DateOfRead,
			ForumModerators:                 moderators,
		}
	}

	response := ForumTopicMessagesDBResponse{
		Topic:                   shortTopic,
		Messages:                messages,
		Attachments:             attachments,
		AdditionalInfos:         additionalInfos,
		TotalMessagesCount:      messageCount,
		MessageDraft:            messageDraft,
		IsUserSubscribed:        isUserSubscribed == 1,
		IsUserTopicAnswerExists: isUserTopicAnswerExists == 1,
		TopicAnswers:            topicAnswers,
		TopicAnsweredUsers:      topicAnsweredUsers,
	}

	return response, nil
}

func (db *DB) FetchForumMessageExists(ctx context.Context, messageId uint64, availableForums []uint64) (bool, error) {
	var isExists uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetMessageExists).WithArgs(messageId, availableForums).FlatArgs(), &isExists)

	if err != nil {
		return false, err
	}

	return isExists == 1, nil
}

func (db *DB) FetchForumMessage(ctx context.Context, messageId uint64) (*ForumMessage, error) {
	var message ForumMessage

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetShortMessage).WithArgs(messageId), &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (db *DB) FetchForumMessageAttachments(ctx context.Context, messageId uint64) ([]ForumMessageAttachment, error) {
	var attachments []ForumMessageAttachment

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessagesAttachments).WithArgs(messageId), &attachments)

	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (db *DB) FetchForumMessageUserVoteExists(ctx context.Context, userId, messageId uint64) (bool, error) {
	var isExists uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetMessageUserVoteExists).WithArgs(userId, messageId), &isExists)

	if err != nil {
		return false, err
	}

	return isExists == 1, nil
}

func (db *DB) FetchUserIsForumModerator(ctx context.Context, userId, forumId uint64) (bool, error) {
	var userIsForumModerator uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetUserIsForumModerator).WithArgs(userId, forumId), &userIsForumModerator)

	if err != nil {
		return false, err
	}

	return userIsForumModerator == 1, nil
}

func (db *DB) InsertForumMessage(ctx context.Context, topic *ForumTopic, userId uint64, login, text string, isRed uint8, forumMessagesInPage uint64) (*ForumMessage, error) {
	var messageId uint64
	var message ForumMessage

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем новое сообщение
				// TODO Поле проставляется для подсчета баллов, однако в методе расчета абсолютно не учитывается нерусский текст
				//  (например, если писать в "English forum", рейтинг никак не поменяется). Выглядит как очередной баг.
				messageLength := len(helpers.RemoveImmeasurable(text))
				result := rw.Write(ctx, sqlapi.NewQuery(queries.ForumInsertNewMessage).WithArgs(messageLength, topic.TopicId, userId, topic.ForumId, isRed, topic.TopicId))
				messageId = uint64(result.LastInsertId)
				return result.Error
			},
			func() error { // Сохраняем текст сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetMessageText).WithArgs(messageId, text)).Error
			},
			func() error { // Получаем сообщение
				return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessages).Inject("").WithArgs(messageId), &message)
			},
			func() error { // Удаляем, если есть, черновик сообщения для данной темы
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumCancelTopicMessagePreview).WithArgs(userId, topic.TopicId)).Error
			},
			func() error { // Обновляем статистику пользователя, выставляем флаг для Cron-а
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumUpdateUserStat).WithArgs(userId)).Error
			},
			func() error { // Обновляем статистику форума
				return updateForumStatAfterNewMessage(ctx, rw, topic, forumMessagesInPage, messageId, userId, login)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	err = notifyForumTopicSubscribersAboutNewMessage(ctx, db.engine, topic.TopicId, messageId, message.DateOfAdd)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func updateForumStatAfterNewMessage(ctx context.Context, rw sqlapi.ReaderWriter, topic *ForumTopic, forumMessagesInPage, messageId, userId uint64, login string) error {
	if topic.Moderated == 0 {
		return nil
	}

	type forumStat struct {
		TopicCount   uint64 `db:"topic_count"`
		MessageCount uint64 `db:"message_count"`
	}
	var stat forumStat
	var topicMessageCount uint64
	var notModeratedTopicCount uint64

	return codeflow.Try(
		func() error { // Обновляем данные о последнем сообщении в теме, выставляем флаг для Cron-а о необходимости пересчета number
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetTopicLastMessage).WithArgs(messageId, userId, login, time.Now(), topic.TopicId)).Error
		},
		func() error { // Выставляем флаг для Cron-а о необходимости переиндексации Sphinx-ом
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMarkTopicNeedSphinxReindex).WithArgs(topic.TopicId)).Error
		},
		func() error { // Получаем обновленную статистику форума
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumStat).WithArgs(topic.ForumId), &stat)
		},
		func() error { // Получаем количество сообщений в теме
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessageCount).WithArgs(topic.TopicId), &topicMessageCount)
		},
		func() error { // Получаем количество непромодерированных тем в форуме
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetNotModeratedTopicCount).WithArgs(topic.ForumId), &notModeratedTopicCount)
		},
		func() error { // Обновляем данные о последней теме в форуме
			pageCount := helpers.CalculatePageCount(topicMessageCount, forumMessagesInPage)
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetForumLastTopic).
				WithArgs(stat.MessageCount, stat.TopicCount, messageId, userId, login, topic.TopicId, topic.Name, time.Now(),
					pageCount, notModeratedTopicCount, topic.ForumId)).Error
		},
	)
}

func notifyForumTopicSubscribersAboutNewMessage(ctx context.Context, rw sqlapi.ReaderWriter, topicId, messageId uint64, messageDate time.Time) error {
	var topicSubscribers []uint64

	type newForumAnswerEntry struct {
		TopicId     uint64    `db:"topic_id"`
		UserId      uint64    `db:"user_id"`
		MessageId   uint64    `db:"message_id"`
		MessageDate time.Time `db:"date_of_add"`
	}

	// Получаем список подписчиков на обновления темы
	err := rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicSubscribers).WithArgs(topicId), &topicSubscribers)

	if err != nil {
		return err
	}

	// Добавляем оповещение для подписчиков о новом сообщении в теме. Не в транзакции, поскольку запрос тяжелый.
	entries := make([]interface{}, 0, len(topicSubscribers))
	for _, userId := range topicSubscribers {
		entries = append(entries, newForumAnswerEntry{
			TopicId:     topicId,
			UserId:      userId,
			MessageId:   messageId,
			MessageDate: messageDate,
		})
	}
	err = rw.Write(ctx, sqlbuilder.InsertInto(queries.NewForumAnswersTable, entries...)).Error

	if err != nil {
		return err
	}

	// Обновляем статистику новых ответов в форуме для подписчиков
	if len(topicSubscribers) > 0 {
		return rw.Write(ctx, sqlapi.NewQuery(queries.ForumIncrementNewForumAnswersCount).WithArgs(topicSubscribers).FlatArgs()).Error
	}

	return nil
}

func (db *DB) FetchForumTopicSubscribers(ctx context.Context, topicId uint64) (subscribers []uint64, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicSubscribers).WithArgs(topicId), &subscribers)
	return
}

func (db *DB) UpdateForumMessage(ctx context.Context, messageId, topicId uint64, text string, isRed uint8) (*ForumMessage, error) {
	var message ForumMessage

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Обновляем сообщение
				// TODO Скорее всего, поле message_length проставляется здесь просто для отчетности. Реального
				//  пересчета баллов, ради которого оно и вводилось, не происходит.
				messageLength := len(helpers.RemoveImmeasurable(text))
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumUpdateMessage).WithArgs(messageLength, isRed, messageId)).Error
			},
			func() error { // Сохраняем текст сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetMessageText).WithArgs(messageId, text)).Error
			},
			func() error { // Получаем сообщение
				return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessages).Inject("").WithArgs(messageId), &message)
			},
			func() error { // Выставляем флаг для Cron-а о необходимости переиндексации Sphinx-ом
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMarkTopicNeedSphinxReindex).WithArgs(topicId)).Error
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (db *DB) DeleteForumMessage(ctx context.Context, messageId, topicId, forumId uint64, messageDate time.Time, forumMessagesInPage uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем сообщение
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteMessage).WithArgs(messageId)).Error
			},
			func() error { // Удаляем текст сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteMessageText).WithArgs(messageId)).Error
			},
			func() error { // Удаляем записи об аттачах сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteMessageFiles).WithArgs(messageId)).Error
			},
			func() error {
				// Записываем сообщение в список удаленных. Таблица чистится при переиндексации форума Sphinx-ом
				// (script/search/source_for_forum_messages.pl)
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMarkMessageDeleted).WithArgs(messageId)).Error
			},
			func() error { // Обновляем данные о последнем сообщении в теме
				return updateTopicStatAfterMessageDeleting(ctx, rw, topicId)
			},
			func() error { // Обновляем данные о последней теме в форуме
				return updateForumStatAfterMessageDeleting(ctx, rw, forumId, forumMessagesInPage)
			},
			func() error { // Помечаем тему, как требующую пересчета (для Cron-а)
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMarkTopicNeedRecalc).WithArgs(topicId)).Error
			},
			func() error {
				// Пересчитываем количество прочитанных пользователями сообщений (используется в списке форумов для
				// подсчета количества непрочитанных сообщений в каждом форуме для данного пользователя)
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumUpdateUserTopicReads).WithArgs(topicId, messageDate)).Error
			},
			func() error {
				return notifyForumTopicSubscribersAboutMessageDeleting(ctx, rw, messageId, topicId)
			},
		)
	})
}

func updateTopicStatAfterMessageDeleting(ctx context.Context, rw sqlapi.ReaderWriter, topicId uint64) error {
	type topicStat struct {
		LastMessageId uint64 `db:"last_message_id"`
	}
	var stat topicStat
	var message ForumMessage

	return codeflow.Try(
		func() error { // Получаем статистику темы
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicStat).WithArgs(topicId), &stat)
		},
		func() error { // Получаем данные о последнем сообщении в теме
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetMessageInfo).WithArgs(stat.LastMessageId), &message)
		},
		func() error { // Обновляем данные о последнем сообщении в теме
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetTopicLastMessage).
				WithArgs(stat.LastMessageId, message.UserId, message.UserLogin, message.DateOfAdd, topicId)).Error
		},
	)
}

func updateForumStatAfterMessageDeleting(ctx context.Context, rw sqlapi.ReaderWriter, forumId, forumMessagesInPage uint64) error {
	type forumStat struct {
		TopicCount   uint64 `db:"topic_count"`
		MessageCount uint64 `db:"message_count"`
	}
	var stat forumStat
	var lastTopic ForumTopic
	var notModeratedTopicCount uint64

	return codeflow.Try(
		func() error { // Получаем статистику форума
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetForumStat).WithArgs(forumId), &stat)
		},
		func() error { // Получаем данные о последней теме в форуме
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetLastTopic).WithArgs(forumId), &lastTopic)
		},
		func() error { // Получаем количество непромодерированных тем в форуме
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetNotModeratedTopicCount).WithArgs(forumId), &notModeratedTopicCount)
		},
		func() error { // Обновляем данные о последней теме в форуме
			pageCount := helpers.CalculatePageCount(lastTopic.MessageCount, forumMessagesInPage)
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumSetForumLastTopic).
				WithArgs(stat.MessageCount, stat.TopicCount, lastTopic.LastMessageId, lastTopic.LastMessageUserId, lastTopic.LastMessageUserLogin,
					lastTopic.TopicId, lastTopic.Name, lastTopic.LastMessageDate, pageCount, notModeratedTopicCount, forumId)).Error
		},
	)
}

func notifyForumTopicSubscribersAboutMessageDeleting(ctx context.Context, rw sqlapi.ReaderWriter, messageId, topicId uint64) error {
	var topicSubscribers []uint64

	return codeflow.Try(
		func() error { // Получаем список подписчиков на обновления темы
			return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicSubscribers).WithArgs(topicId), &topicSubscribers)
		},
		func() error { // Удаляем оповещение о новом сообщении, если есть, для всех подписчиков
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteNewForumAnswer).WithArgs(messageId)).Error
		},
		func() error { // Обновляем статистику новых ответов в форуме для подписчиков
			return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDecrementNewForumAnswersCount).WithArgs(topicSubscribers).FlatArgs()).Error
		},
	)
}
