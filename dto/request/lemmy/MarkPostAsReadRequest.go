package lemmy

// Read is deliberately NOT validate:"required" — false is a legitimate
// value here (marking a post unread), and go-playground/validator's
// "required" tag rejects the zero value on plain bool fields, which would
// silently break every "mark as unread" request.
// One of PostId or PostIds should be set.
type MarkPostAsReadRequest struct {
	PostId  *uint  `json:"post_id,omitempty"`
	PostIds []uint `json:"post_ids,omitempty"`
	Read    bool   `json:"read"`
}
