package piefed

// Read is deliberately NOT validate:"required" — see the lemmy-side
// MarkPostAsReadRequest for why (false/unread is a legitimate value).
type MarkPostAsReadRequest struct {
	PostId  *uint  `json:"post_id,omitempty"`
	PostIds []uint `json:"post_ids,omitempty"`
	Read    bool   `json:"read"`
}
