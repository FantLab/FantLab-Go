package forumapi

// Wrapper для списка форумов
type forumBlocksWrapper struct {
	ForumBlocks []forumBlock `json:"forum_blocks"`
}

// Блок форумов
type forumBlock struct {
	ID     uint16  `json:"-"`
	Title  string  `json:"block_title"`
	Forums []forum `json:"forums"`
}

// Форум
type forum struct {
	ID          uint16      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Moderators  []userLink  `json:"moderators"`
	Stats       forumStats  `json:"stats"`
	LastMessage lastMessage `json:"last_message"`
}

// Wrapper для списке тем
type forumTopicsWrapper struct {
	Topics []forumTopic `json:"topics"`
}

// Тема
type forumTopic struct {
	ID          uint32      `json:"id"`
	Title       string      `json:"title"`
	TopicType   uint16      `json:"topic_type"`
	Creation    creation    `json:"creation"`
	IsClosed    bool        `json:"is_closed"`
	IsPinned    bool        `json:"is_pinned"`
	Stats       topicStats  `json:"stats"`
	LastMessage lastMessage `json:"last_message"`
}

// Wrapper для списка сообщений
type topicMessagesWrapper struct {
	Messages []topicMessage `json:"messages"`
}

// Сообщение в форуме
// IsCensored - сообщение изъято модератором
// IsModerTagWorks - рендерить ли в сообщении содержимое тега [moder] (true, только если тег добавлен модератором)
type topicMessage struct {
	ID              uint32       `json:"id"`
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
	ID    uint32     `json:"id"`
	Topic *topicLink `json:"topic,omitempty"`
	User  userLink   `json:"user"`
	Date  int64      `json:"date"`
}

// Ссылка на тему форума
type topicLink struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

// Ссылка на пользователя
// PhotoNumber - порядковый номер фото (https://data.fantlab.ru/images/users/{UserId}_{PhotoNumber}). Если 0 - его нет.
type userLink struct {
	ID          uint32 `json:"id"`
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
