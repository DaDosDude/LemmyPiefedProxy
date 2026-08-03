package lemmy

type EditPostRequest struct {
	PostId     uint    `json:"post_id" validate:"required"`
	Name       *string `json:"name,omitempty"`
	Body       *string `json:"body,omitempty"`
	Url        *string `json:"url,omitempty"`
	Nsfw       *bool   `json:"nsfw,omitempty"`
	LanguageId *uint   `json:"language_id,omitempty"`
}
