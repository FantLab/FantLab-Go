package shared

type AppConfig struct {
	ImagesBaseURL         string
	BlogsInPage           uint16
	BlogTopicsInPage      uint16
	ForumTopicsInPage     uint32
	ForumMessagesInPage   uint32
	DefaultAccessToForums []uint16
}
