package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunityView(in piefed.CommunityView) lemmy.CommunityView {
	return lemmy.CommunityView{
		BannedFromCommunity: in.BannedFromCommunity,
		Blocked:             in.Blocked,
		Community:           ConvertCommunity(in.Community),
		Counts:              ConvertCommunityAggregates(in.Counts),
		Subscribed:          ConvertSubscribedType(in.Subscribed),
	}
}
