package config

import "time"

const (
	AuthTokenTimeout    = 10 * time.Minute
	RefreshTokenTimeout = 30 * 24 * time.Hour
)

type AppConfig struct {
	SiteUrl                                  string
	SiteName                                 string
	SiteEmail                                string
	BaseImageUrl                             string
	BaseMinioFileUrl                         string
	BaseForumMessageAttachUrl                string
	BaseForumMessageDraftAttachUrl           string
	MinResponseLength                        uint64
	MaxUserResponseCountPerWork              uint64
	MinUserOwnResponsesRatingForMinusAbility uint64
	ForumTopicsInPage                        uint64
	ForumMessagesInPage                      uint64
	MaxMessageLength                         uint64
	MaxForumMessageEditTimeout               uint64
	DefaultAccessToForums                    []uint64
	ForumsWithEnabledRating                  []uint64
	ForumsWithDisabledMinuses                []uint64
	ReadOnlyForumUsers                       map[uint64][]uint64
	BlogsInPage                              uint64
	BlogTopicsInPage                         uint64
	BlogArticleCommentsInPage                uint64
	CensorshipText                           string
	PreModerationText                        string
	ImageReplacementLinkText                 string
	BotUserId                                uint64
	MaxAttachCountPerMessage                 uint64
	BookcaseItemInPage                       uint64
	Smiles                                   *Smiles
	FlContestInProgress                      bool
	FlContestAuthorId                        uint64
	CorrelationUserMarkCountThreshold        uint64
}
