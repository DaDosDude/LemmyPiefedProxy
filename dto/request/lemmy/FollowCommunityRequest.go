package lemmy

// Follow is deliberately NOT validate:"required" — false is a legitimate
// value here (leaving a community), and go-playground/validator's
// "required" tag rejects the zero value on plain bool fields.
type FollowCommunityRequest struct {
	CommunityId uint `json:"community_id" validate:"required"`
	Follow      bool `json:"follow"`
}
