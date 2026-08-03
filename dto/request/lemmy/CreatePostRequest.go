package lemmy

// Name is Lemmy's field for what Piefed calls Title. Honeypot (an
// anti-spam field with no content of its own) is accepted here so the
// request doesn't fail parsing, but has nothing to forward to and is
// silently dropped.
type CreatePostRequest struct {
	Name        string  `json:"name" validate:"required"`
	CommunityId uint    `json:"community_id" validate:"required"`
	Body        *string `json:"body,omitempty"`
	Url         *string `json:"url,omitempty"`
	Nsfw        *bool   `json:"nsfw,omitempty"`
	LanguageId  *uint   `json:"language_id,omitempty"`
	Honeypot    *string `json:"honeypot,omitempty"`
}
