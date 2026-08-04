package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
	"LemmyBeProxy/helper"
)

func ConvertPostView(in piefed.PostView) lemmy.PostView {
	return lemmy.PostView{
		BannedFromCommunity:        in.BannedFromCommunity,
		Community:                  ConvertCommunity(in.Community),
		Counts:                     ConvertPostAggregates(in.Counts),
		Creator:                    ConvertPerson(in.Creator),
		CreatorBannedFromCommunity: in.CreatorBannedFromCommunity,
		CreatorBlocked:             helper.DefaultOnNil(in.CreatorBlocked, false),
		CreatorIsAdmin:             in.CreatorIsAdmin,
		CreatorIsModerator:         in.CreatorIsModerator,
		Hidden:                     in.Hidden,
		ImageDetails:               nil,
		MyVote:                     in.MyVote,
		Post:                       ConvertPost(in.Post),
		Read:                       in.Read,
		Saved:                      in.Saved,
		Subscribed:                 ConvertSubscribedType(in.Subscribed),
		UnreadComments:             in.UnreadComments,
	}
}
