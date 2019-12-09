package shared

type AppConfig struct {
	ImagesBaseURL         string
	BlogsInPage           uint64
	BlogTopicsInPage      uint64
	ForumTopicsInPage     uint64
	ForumMessagesInPage   uint64
	DefaultAccessToForums []uint64
	CensorshipText        string
}
