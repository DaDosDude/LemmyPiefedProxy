package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommentAggregates(in piefed.CommentAggregates) lemmy.CommentAggregates {
	return lemmy.CommentAggregates{
		ChildCount: in.ChildCount,
		CommentId:  in.CommentId,
		Downvotes:  in.Downvotes,
		Published:  in.Published,
		Score:      in.Score,
		Upvotes:    in.Upvotes,
	}
}
