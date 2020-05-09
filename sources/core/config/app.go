package config

import "time"

const (
	AuthTokenTimeout    = 10 * time.Minute
	RefreshTokenTimeout = 30 * 24 * time.Hour
)

type AppConfig struct {
	SiteURL                                  string
	SiteName                                 string
	SiteEmail                                string
	ImagesBaseURL                            string
	MinUserOwnResponsesRatingForMinusAbility uint64
	ForumTopicsInPage                        uint64
	ForumMessagesInPage                      uint64
	MaxForumMessageLength                    uint64
	MaxForumMessageEditTimeout               uint64
	DefaultAccessToForums                    []uint64
	BlogsInPage                              uint64
	BlogTopicsInPage                         uint64
	BlogArticleCommentsInPage                uint64
	CensorshipText                           string
	BotUserId                                uint64
	MaxAttachCountPerMessage                 uint64
	BookcaseItemInPage                       uint64
}
