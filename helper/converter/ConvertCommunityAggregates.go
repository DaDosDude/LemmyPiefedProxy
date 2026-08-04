package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunityAggregates(in piefed.CommunityAggregates) lemmy.CommunityAggregates {
	return lemmy.CommunityAggregates{
		Comments:            in.PostReplyCount,
		CommunityId:         in.CommunityId,
		Posts:               in.PostCount,
		Published:           in.Published,
		Subscribers:         in.SubscriptionsCount,
		SubscribersLocal:    0,
		UsersActiveDay:      in.ActiveDaily,
		UsersActiveHalfYear: in.ActiveSixMonthly,
		UsersActiveMonth:    in.ActiveMonthly,
		UsersActiveWeek:     in.ActiveWeekly,
	}
}
