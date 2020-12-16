package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Статьи в блогах

func GetBlogArticleAttachmentsDir(articleId uint64) string {
	return fmt.Sprintf("/blog_files/b%d", articleId)
}

func GetBlogArticleAttachments(articleId uint64) ([]File, error) {
	var attachments []File

	files, err := ioutil.ReadDir(GetBlogArticleAttachmentsDir(articleId))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()
		if fileName != "img" { // пропускаем директорию с картинками внутри текста
			attachments = append(attachments, File{
				Name: fileName,
				Size: uint64(file.Size()),
			})
		}
	}

	return attachments, nil
}

// Сообщения в форуме

func GetForumMessageAttachmentsDir(messageId uint64) string {
	return "/forum_attach/" + IdToRelativeFilePath(messageId, 3)
}

// Метода GetForumMessageAttachments нет, поскольку список аттачей к сообщению форума вытаскивается из базы. Разумеется,
// это быстрее, но чревато ошибками из-за рассинхрона с реальным содержимым файловой системы (например, после факапа,
// влекущего за собой удаление части файлов).

func DeleteForumMessageAttachment(messageId uint64, name string) {
	fileName := fmt.Sprintf("%s/%s", GetForumMessageAttachmentsDir(messageId), name)
	_ = os.Remove(fileName)
}

func DeleteForumMessageAttachments(messageId uint64) {
	_ = os.RemoveAll(GetForumMessageAttachmentsDir(messageId))
}

// Черновики сообщений в форуме

func GetForumMessageDraftAttachmentsDir(userId, topicId uint64) string {
	return fmt.Sprintf("/files/preview/m_%d_%d", userId, topicId)
}

func GetForumMessageDraftAttachments(userId, topicId uint64) ([]File, error) {
	var attachments []File

	files, err := ioutil.ReadDir(GetForumMessageDraftAttachmentsDir(userId, topicId))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		attachments = append(attachments, File{
			Name: file.Name(),
			Size: uint64(file.Size()),
		})
	}

	return attachments, nil
}

func DeleteForumMessageDraftAttachment(userId, topicId uint64, name string) {
	fileName := fmt.Sprintf("%s/%s", GetForumMessageDraftAttachmentsDir(userId, topicId), name)
	_ = os.Remove(fileName)
}

func DeleteForumMessageDraftAttachments(userId, topicId uint64) {
	_ = os.RemoveAll(GetForumMessageDraftAttachmentsDir(userId, topicId))
}
