package queries

const (
	ForumCancelTopicMessagePreview = `
		DELETE FROM
			f_messages_preview
		WHERE
			user_id = ? AND topic_id = ?
	`

	ForumInsertMessagePreview = `
		INSERT INTO
			f_messages_preview (
				message,
				user_id,
				topic_id,
				date_of_add,
				date_of_edit
			)
		VALUES
			(?, ?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE
			message = ?,
			date_of_edit = NOW()
	`

	ForumGetTopicMessagePreview = `
		SELECT
			f.preview_id,
			f.topic_id,
			f.message,
			f.date_of_add,
			f.date_of_edit,
			f.user_id,
			u.login,
			u.sex,
			u.photo_number,
			u.user_class,
			u.sign
		FROM
			f_messages_preview f
		LEFT JOIN
			users u ON u.user_id = f.user_id
		WHERE
			f.topic_id = ? AND f.user_id = ?
		LIMIT 1
	`

	ForumDeleteForumMessagePreview = `
		DELETE
		FROM
			f_messages_preview
		WHERE
			topic_id = ? AND user_id = ?
	`

	ForumGetMessageDraftMinioFileCount = `
		SELECT
			COUNT(*)
		FROM
			files_minio
		WHERE
			file_group = 'forum_draft' AND holder_id = ?
	`

	ForumGetMessageDraftMinioFiles = `
		SELECT
			file_id,
			file_group,
			holder_id AS message_id,
			file_name,
			file_size,
			date_of_add,
			user_id
		FROM
			files_minio
		WHERE
			file_group = 'forum_draft' AND holder_id = ?
	`

	ForumGetMessageDraftMinioFile = `
		SELECT
			file_id,
			file_group,
			holder_id AS draft_id,
			file_name,
			file_size,
			date_of_add,
			user_id
		FROM
			files_minio
		WHERE
			file_group = 'forum_draft' AND holder_id = ? AND file_id = ?
	`

	// Ignore on duplicate (very unlikely): https://stackoverflow.com/a/4596409
	ForumInsertMessageDraftMinioFile = `
		INSERT INTO
			files_minio (
				file_group,
				holder_id,
				file_name,
				file_size,
				date_of_add,
				user_id
			)
		VALUES
			('forum_draft', ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			file_id = file_id
	`

	ForumDeleteMessageDraftMinioFile = `
		DELETE
		FROM
			files_minio
		WHERE
			file_group = 'forum_draft' AND holder_id = ? AND file_id = ?
	`
)
