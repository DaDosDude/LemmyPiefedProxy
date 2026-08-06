package lemmy017

// CommunityAggregates matches Lemmy 0.17.2's real CommunityAggregates.
// Same extra-id pattern as PostAggregates/CommentAggregates — Id
// defaulted to CommunityId's value.
type CommunityAggregates struct {
	Id                  uint   `json:"id" validate:"required"`
	CommunityId         uint   `json:"community_id" validate:"required"`
	Subscribers         uint   `json:"subscribers" validate:"required"`
	Posts               uint   `json:"posts" validate:"required"`
	Comments            uint   `json:"comments" validate:"required"`
	Published           string `json:"published" validate:"required"`
	UsersActiveDay      uint   `json:"users_active_day" validate:"required"`
	UsersActiveWeek     uint   `json:"users_active_week" validate:"required"`
	UsersActiveMonth    uint   `json:"users_active_month" validate:"required"`
	UsersActiveHalfYear uint   `json:"users_active_half_year" validate:"required"`
}
