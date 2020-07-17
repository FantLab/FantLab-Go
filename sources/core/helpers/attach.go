package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetBlogArticleAttachments(articleId uint64) ([]File, error) {
	var attachments []File

	files, err := ioutil.ReadDir(fmt.Sprintf("/blog_files/b%d", articleId))
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

func DeleteForumMessageAttachments(messageId uint64) {
	_ = os.RemoveAll("/forum_attach/" + IdToRelativeFilePath(messageId, 3))
}

func DeleteForumMessageDraftAttachments(userId, topicId uint64) {
	_ = os.RemoveAll(fmt.Sprintf("/files/preview/m_%d_%d", userId, topicId))
}
