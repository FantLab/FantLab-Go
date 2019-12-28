package queries

const (
	SessionsTable = "sessions2" // sessions https://github.com/parserpro/fantlab/issues/908
	UsersTable    = "users"     // users2
)

const (
	UserSession       = "SELECT user_id, date_of_create FROM " + SessionsTable + " WHERE code = ? LIMIT 1"
	UserPasswordHash  = "SELECT user_id, password_hash, new_password_hash FROM " + UsersTable + " WHERE login = ? LIMIT 1"
	UserBlock         = "SELECT block, date_of_block_end, block_reason FROM " + UsersTable + " WHERE user_id = ? LIMIT 1"
	UserClass         = "SELECT user_class from " + UsersTable + " WHERE user_id = ? LIMIT 1"
	DeleteUserSession = "DELETE FROM " + SessionsTable + " WHERE code = ?"
)
