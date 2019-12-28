package queries

const (
	FetchBlogTopicLikeCount   = "SELECT likes_count FROM b_topics WHERE topic_id = ?"
	IsBlogTopicLiked          = "SELECT 1 FROM b_topic_likes WHERE topic_id = ? AND user_id = ?"
	LikeBlogTopic             = "INSERT INTO b_topic_likes (topic_id, user_id, date_of_add) VALUES (?, ?, ?)"
	DislikeBlogTopic          = "DELETE FROM b_topic_likes WHERE topic_id = ? AND user_id = ?"
	UpdateBlogTopicLikesCount = `
		UPDATE b_topics b
		SET b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)
		WHERE b.topic_id = ?
	`
)
