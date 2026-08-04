package lemmy017

type CreatePostRequest struct {
	Name        string  `json:"name" validate:"required"`
	CommunityId uint    `json:"community_id" validate:"required"`
	Url         *string `json:"url,omitempty"`
	Body        *string `json:"body,omitempty"`
	Honeypot    *string `json:"honeypot,omitempty"`
	Nsfw        *bool   `json:"nsfw,omitempty"`
	LanguageId  *uint   `json:"language_id,omitempty"`
}
