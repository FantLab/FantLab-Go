package app

import (
	"fantlab/core/helpers"
	"fmt"
)

func (s *Services) GetFSForumMessageAttachmentUrl(messageId uint64, fileName string) string {
	return fmt.Sprintf("%s/%s/%s", s.appConfig.BaseForumMessageAttachUrl, helpers.IdToRelativeFilePath(messageId, 3), fileName)
}

func (s *Services) GetFSForumMessageDraftAttachmentUrl(userId, topicId uint64, fileName string) string {
	return fmt.Sprintf("%s/%s/%s", s.appConfig.BaseForumMessageDraftAttachUrl, fmt.Sprintf("m_%d_%d", userId, topicId), fileName)
}

func (s *Services) GetMinioForumMessageAttachmentUrl(messageId uint64, fileName string) string {
	return fmt.Sprintf("%s/%s/%d/%s", s.appConfig.BaseMinioFileUrl, ForumMessageFileGroup, messageId, fileName)
}

func (s *Services) GetMinioForumMessageDraftAttachmentUrl(messageDraftId uint64, fileName string) string {
	return fmt.Sprintf("%s/%s/%d/%s", s.appConfig.BaseMinioFileUrl, ForumMessageDraftFileGroup, messageDraftId, fileName)
}
