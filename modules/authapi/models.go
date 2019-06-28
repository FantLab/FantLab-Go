package authapi

type session struct {
	UserId  uint32 `json:"user_id"`
	Session string `json:"session"`
}
