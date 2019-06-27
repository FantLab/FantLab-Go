package authapi

type session struct {
	UserId  int    `json:"user_id"`
	Session string `json:"session"`
}
