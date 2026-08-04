package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetSiteResponse struct {
	Admins              []lemmy.PersonView            `json:"admins" validate:"required"`
	AllLanguages        []lemmy.Language              `json:"all_languages" validate:"required"`
	BlockedUrls         []lemmy.LocalSiteUrlBlocklist `json:"blocked_urls" validate:"required"`
	CustomEmojis        []lemmy.CustomEmojiView       `json:"custom_emojis" validate:"required"`
	DiscussionLanguages []uint                        `json:"discussion_languages" validate:"required"`
	MyUser              *lemmy.MyUserInfo             `json:"my_user,omitempty"`
	SiteView            lemmy.SiteView                `json:"site_view" validate:"required"`
	Taglines            []lemmy.Tagline               `json:"taglines" validate:"required"`
	Version             string                        `json:"version" validate:"required"`
}
