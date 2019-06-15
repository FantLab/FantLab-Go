package blogsapi

// Wrapper для списка рубрик
type communitiesWrapper struct {
	Main       []community `json:"main"`
	Additional []community `json:"additional"`
}

// Рубрика
type community struct {
	Id          uint32               `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Stats       blogStats            `json:"stats"`
	LastArticle lastCommunityArticle `json:"last_article"`
}

// Статистика блога
type blogStats struct {
	ArticleCount    uint32 `json:"article_count"`
	SubscriberCount uint32 `json:"subscriber_count"`
}

// Последняя статья в рубрике
type lastCommunityArticle struct {
	Id    uint32   `json:"id"`
	Title string   `json:"title"`
	User  userLink `json:"user"`
	Date  int64    `json:"date"`
}

// Ссылка на пользователя
type userLink struct {
	Id    uint32 `json:"id"`
	Login string `json:"login"`
}
