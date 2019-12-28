package queries

const (
	BlogSubscriptionExists = "SELECT 1 FROM b_subscribers WHERE blog_id = ? AND user_id = ?"
	BlogSubscriptionInsert = "INSERT INTO b_subscribers (user_id, blog_id, date_of_add) VALUES (?, ?, ?)"
	BlogSubscriptionDelete = "DELETE FROM b_subscribers WHERE blog_id = ? AND user_id = ?"
	BlogSubscriberUpdate   = `
		UPDATE b_blogs b
		SET b.subscriber_count = (SELECT COUNT(DISTINCT bs.user_id) FROM b_subscribers bs WHERE bs.blog_id = b.blog_id)
		WHERE b.blog_id = ?
	`
	BlogTopicSubscriptionExists = "SELECT 1 FROM b_topics_subscribers WHERE topic_id = ? AND user_id = ?"
	BlogTopicSubscriptionInsert = "INSERT INTO b_topics_subscribers (user_id, topic_id, date_of_add) VALUES (?, ?, ?)"
	BlogTopicSubscriptionDelete = "DELETE FROM b_topics_subscribers WHERE topic_id = ? AND user_id = ?"
)
