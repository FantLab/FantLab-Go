package helpers

import (
	"fmt"
	"os"
)

func DeleteForumMessageTextCache(messageId uint64) {
	_ = os.Remove("/cache/f_messages/" + IdToRelativeFilePath(messageId, 3))
}

func DeleteBlogCommentTextCache(commentId uint64) {
	_ = os.Remove("/cache/blog_comments/" + IdToRelativeFilePath(commentId, 2))
}

func DeleteResponseTextCache(responseId uint64) {
	_ = os.Remove("/cache/responses/" + IdToRelativeFilePath(responseId, 3))
}

func DeleteWorkRatingImageCache(workId uint64) {
	_ = os.Remove(fmt.Sprintf("/cache_img/rating/%d.gif", workId))
}
