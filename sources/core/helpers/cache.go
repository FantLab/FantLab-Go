package helpers

import "os"

func DeleteForumMessageTextCache(messageId uint64) {
	_ = os.Remove("/cache/f_messages/" + IdToRelativeFilePath(messageId, 3))
}

func DeleteBlogCommentTextCache(commentId uint64) {
	_ = os.Remove("/cache/blog_comments/" + IdToRelativeFilePath(commentId, 2))
}
