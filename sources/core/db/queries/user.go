package queries

const (
	UsersTable      = "users"
	AuthTokensTable = "auth_tokens"
)

const (
	UserLoginInfo = "SELECT user_id, password_hash, new_password_hash, approved FROM " + UsersTable + " WHERE login = ? OR email = ? LIMIT 1"

	UserInfo = `
		SELECT 
			u.user_class,
			u.login,
			u.sex,
			u.votecount,
			g.can_edit_delete_f_messages,
			g.can_edit_f_messages,
			g.access_to_forums,
			g.can_edit_responses,
			s.always_pm_by_email
		FROM
			users u
		LEFT JOIN
			user_groups g ON g.user_group_id = u.user_group_id
		LEFT JOIN
			user_settings s ON s.user_id = u.user_id
		WHERE
			u.user_id = ?
		LIMIT 1
	`

	UserBlock          = "SELECT block, date_of_block_end, block_reason FROM " + UsersTable + " WHERE user_id = ? LIMIT 1"
	FetchAuthTokenById = "SELECT * FROM " + AuthTokensTable + " WHERE token_id = ? LIMIT 1"
	DeleteAuthToken    = "DELETE FROM " + AuthTokensTable + " WHERE token_id = ?"

	UserGetInfo = `
		SELECT
			user_id,
			login,
			email
		FROM
			users
		WHERE
			user_id = ?
	`

	UserGetMarkCount = `
		SELECT
			markcount
		FROM
			users
		WHERE
			user_id = ?
	`

	UserMarkUserNeedLevelRecalc = `
		UPDATE
			users
		SET
			need_recalc_level = 1
		WHERE
			user_id = ?
	`

	UserIncrementResponseCount = `
		UPDATE
			users
		SET
			responsecount = responsecount + 1
		WHERE
			user_id = ?
	`

	UserDecrementResponseCount = `
		UPDATE
			users
		SET
			responsecount = responsecount - 1
		WHERE
			user_id = ?
	`

	UserUpdateMarkCount = `
		UPDATE
			users u
		SET
			u.markcount = (SELECT COUNT(DISTINCT m.work_id) FROM marks2 m WHERE m.user_id = u.user_id),
			u.need_recalc_level = 1
		WHERE
			u.user_id = ?
	`
)
