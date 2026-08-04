package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommentView(in piefed.CommentView) lemmy.CommentView {
	return lemmy.CommentView{
		BannedFromCommunity:        in.BannedFromCommunity,
		Comment:                    ConvertComment(in.Comment),
		Community:                  ConvertCommunity(in.Community),
		Counts:                     ConvertCommentAggregates(in.Counts),
		Creator:                    ConvertPerson(in.Creator),
		CreatorBannedFromCommunity: in.CreatorBannedFromCommunity,
		CreatorBlocked:             in.CreatorBlocked,
		CreatorIsAdmin:             in.CreatorIsAdmin,
		CreatorIsModerator:         in.CreatorIsModerator,
		MyVote:                     in.MyVote,
		Post:                       ConvertPost(in.Post),
		Saved:                      in.Saved,
		Subscribed:                 ConvertSubscribedType(in.Subscribed),
	}
}
