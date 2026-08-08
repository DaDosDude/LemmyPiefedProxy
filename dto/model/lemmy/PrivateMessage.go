package lemmy

type PrivateMessage struct {
	Id          uint    `json:"id" validate:"required"`
	CreatorId   uint    `json:"creator_id" validate:"required"`
	RecipientId uint    `json:"recipient_id" validate:"required"`
	Content     string  `json:"content" validate:"required"`
	Deleted     bool    `json:"deleted"`
	Read        bool    `json:"read"`
	Published   string  `json:"published" validate:"required"`
	Updated     *string `json:"updated,omitempty"`
	ApId        string  `json:"ap_id" validate:"required"`
	Local       bool    `json:"local"`
}
