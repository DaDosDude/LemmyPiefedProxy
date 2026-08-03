package piefed

type CreatePostRequest struct {
	Title       string  `json:"title" validate:"required"`
	CommunityId uint    `json:"community_id" validate:"required"`
	Body        *string `json:"body,omitempty"`
	Url         *string `json:"url,omitempty"`
	Nsfw        *bool   `json:"nsfw,omitempty"`
	LanguageId  *uint   `json:"language_id,omitempty"`
}
