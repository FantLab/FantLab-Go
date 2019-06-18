package blogsapi

// Wrapper для списка рубрик
type communitiesWrapper struct {
	Main       []community `json:"main"`
	Additional []community `json:"additional"`
}

// Рубрика
type community struct {
	Id          uint32      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Stats       stats       `json:"stats"`
	LastArticle lastArticle `json:"last_article"`
}

// Wrapper для списка авторских колонок
type blogsWrapper struct {
	Blogs []blog `json:"blogs"`
}

// Авторская колонка
type blog struct {
	Id          uint32      `json:"id"`
	Owner       userLink    `json:"owner"`
	IsClosed    bool        `json:"is_closed"`
	Stats       stats       `json:"stats"`
	LastArticle lastArticle `json:"last_article"`
}

// Статистика
type stats struct {
	ArticleCount    uint32 `json:"article_count"`
	SubscriberCount uint32 `json:"subscriber_count"`
}

// Последняя статья
type lastArticle struct {
	Id    uint32    `json:"id"`
	Title string    `json:"title"`
	User  *userLink `json:"user,omitempty"`
	Date  int64     `json:"date"`
}

// Ссылка на пользователя
type userLink struct {
	Id    uint32 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name,omitempty"`
}
