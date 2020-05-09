package queries

const (
	PrivateGetMessage = `
		SELECT
			pm.private_message_id,
			pm.from_user_id,
			pm.to_user_id,
			pm.is_read,
			pm.date_of_add,
			u.login,
			u.sex,
			u.photo_number,
			u.user_class,
			u.sign,
			t.message_text,
			pm.number
		FROM
			f_private_messages pm
		LEFT JOIN
			users u ON u.user_id = pm.from_user_id
		LEFT JOIN
			f_private_messages_text t ON t.private_message_id = pm.private_message_id
		WHERE
			pm.private_message_id = ?
		LIMIT 1
	`

	PrivateInsertNewMessage = `
		INSERT INTO
			f_private_messages (
				from_user_id,
				to_user_id,
				copy_by_email,
				date_of_add,
				is_red,
				number
			)
		SELECT
			?, ?, ?, NOW(), ?, COALESCE(MAX(number), 0) + 1
		FROM
			f_private_messages
		WHERE
			(from_user_id = ? AND to_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)
	`

	PrivateInsertMessageText = `
		INSERT INTO
			f_private_messages_text (
				private_message_id,
				message_text
			)
		VALUES
			(?, ?)
	`

	PrivateCancelMessagePreview = `
		DELETE
		FROM
			f_private_messages_preview
		WHERE
			from_user_id = ? AND to_user_id = ?
	`
)
