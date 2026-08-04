package lemmy017

type EditPostRequest struct {
	PostId     uint    `json:"post_id" validate:"required"`
	Name       *string `json:"name,omitempty"`
	Url        *string `json:"url,omitempty"`
	Body       *string `json:"body,omitempty"`
	Nsfw       *bool   `json:"nsfw,omitempty"`
	LanguageId *uint   `json:"language_id,omitempty"`
}
