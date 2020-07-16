package helpers

import (
	"fmt"
	"os"
)

func DeleteForumMessageAttachments(messageId uint64) {
	_ = os.RemoveAll("/forum_attach/" + IdToRelativeFilePath(messageId, 3))
}

func DeleteForumMessageDraftAttachments(userId, topicId uint64) {
	_ = os.RemoveAll(fmt.Sprintf("/files/preview/m_%d_%d", userId, topicId))
}
