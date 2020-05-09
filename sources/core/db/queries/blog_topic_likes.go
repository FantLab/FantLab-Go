package queries

const (
	FetchBlogTopicLikeCount = "SELECT likes_count FROM b_topics WHERE topic_id = ?"
	LikeBlogTopic           = `
		INSERT INTO b_topic_likes (topic_id, user_id, date_of_add) VALUES (?, ?, NOW())
		ON DUPLICATE KEY UPDATE date_of_add = NOW()
	`
	DislikeBlogTopic          = "DELETE FROM b_topic_likes WHERE topic_id = ? AND user_id = ?"
	UpdateBlogTopicLikesCount = `
		UPDATE b_topics b
		SET b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)
		WHERE b.topic_id = ?
	`
)
