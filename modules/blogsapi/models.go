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
	Description string      `json:"description,omitempty"`
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
	User        userLink    `json:"user"`
	IsClosed    bool        `json:"is_closed"`
	Stats       stats       `json:"stats"`
	LastArticle lastArticle `json:"last_article"`
}

// Wrapper для списка статей в авторской колонке
type blogArticlesWrapper struct {
	Articles []article `json:"articles"`
}

// Статья
type article struct {
	Id       uint32       `json:"id"`
	Title    string       `json:"title"`
	Creation creation     `json:"creation"`
	Text     string       `json:"text"`
	Tags     string       `json:"tags"`
	Stats    articleStats `json:"stats"`
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
	Id     uint32 `json:"id"`
	Login  string `json:"login"`
	Name   string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// Данные о создании
type creation struct {
	User userLink `json:"user"`
	Date int64    `json:"date"`
}

// Статистика статьи
type articleStats struct {
	LikeCount    uint64 `json:"like_count"`
	ViewCount    uint32 `json:"view_count"`
	CommentCount uint32 `json:"comment_count"`
}
