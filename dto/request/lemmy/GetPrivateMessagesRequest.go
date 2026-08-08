package lemmy

type GetPrivateMessagesRequest struct {
	UnreadOnly *bool `json:"unread_only,omitempty"`
	Page       *uint `json:"page,omitempty"`
	Limit      *uint `json:"limit,omitempty"`
}
