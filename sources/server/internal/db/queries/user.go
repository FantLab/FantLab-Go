package queries

const (
	UsersTable      = "users"
	UserGroupsTable = "user_groups"
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
			g.can_edit_f_messages,
			g.access_to_forums,
			g.can_edit_responses
		FROM ` + UserGroupsTable + ` g
		JOIN ` + UsersTable + ` u ON u.user_group_id = g.user_group_id
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

	UserMarkUserNeedLevelRecalc = `
		UPDATE
			users
		SET
			need_recalc_level = 1
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
)
