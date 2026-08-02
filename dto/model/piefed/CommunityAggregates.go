package piefed

type CommunityAggregates struct {
	CommunityId        uint   `json:"id" validate:"required"`
	SubscriptionsCount uint   `json:"subscriptions_count" validate:"required"`
	PostCount          uint   `json:"post_count" validate:"required"`
	PostReplyCount     uint   `json:"post_reply_count" validate:"required"`
	Published          string `json:"published" validate:"required"`
	ActiveDaily        uint   `json:"active_daily" validate:"required"`
	ActiveWeekly       uint   `json:"active_weekly" validate:"required"`
	ActiveMonthly      uint   `json:"active_monthly" validate:"required"`
	ActiveSixMonthly   uint   `json:"active_6monthly" validate:"required"`
}
