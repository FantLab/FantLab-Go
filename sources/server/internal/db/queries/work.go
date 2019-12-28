package queries

const (
	WorkExists   = "SELECT 1 FROM works WHERE work_id = ?"
	WorkUserMark = "SELECT mark FROM marks2 WHERE user_id = ? AND work_id = ?"
)
