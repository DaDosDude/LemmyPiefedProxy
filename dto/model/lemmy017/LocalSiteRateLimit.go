package lemmy017

// LocalSiteRateLimit matches Lemmy 0.17.2's real LocalSiteRateLimit.
// Same extra-id pattern as other Aggregates-adjacent types.
// ImportUserSettings/ImportUserSettingsPerSecond are a later addition
// our canonical model has but 0.17.2 doesn't — simply not present here.
type LocalSiteRateLimit struct {
	Id                uint    `json:"id" validate:"required"`
	LocalSiteId       uint    `json:"local_site_id" validate:"required"`
	Message           uint    `json:"message" validate:"required"`
	MessagePerSecond  uint    `json:"message_per_second" validate:"required"`
	Post              uint    `json:"post" validate:"required"`
	PostPerSecond     uint    `json:"post_per_second" validate:"required"`
	Register          uint    `json:"register" validate:"required"`
	RegisterPerSecond uint    `json:"register_per_second" validate:"required"`
	Image             uint    `json:"image" validate:"required"`
	ImagePerSecond    uint    `json:"image_per_second" validate:"required"`
	Comment           uint    `json:"comment" validate:"required"`
	CommentPerSecond  uint    `json:"comment_per_second" validate:"required"`
	Search            uint    `json:"search" validate:"required"`
	SearchPerSecond   uint    `json:"search_per_second" validate:"required"`
	Published         string  `json:"published" validate:"required"`
	Updated           *string `json:"updated"`
}
