package lemmy

// Block is deliberately NOT validate:"required" — false is a legitimate
// value here (unblocking), and go-playground/validator's "required" tag
// rejects the zero value on plain bool fields.
type BlockCommunityRequest struct {
	CommunityId uint `json:"community_id" validate:"required"`
	Block       bool `json:"block"`
}
