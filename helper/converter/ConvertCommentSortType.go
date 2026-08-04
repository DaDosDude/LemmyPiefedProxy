package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommentSortType(in piefed.CommentSortType) lemmy.CommentSortType {
	return lemmy.CommentSortType(in)
}

func ReverseConvertCommentSortType(in lemmy.CommentSortType) piefed.CommentSortType {
	if in == lemmy.CommentSortTypeControversial {
		return piefed.CommentSortTypeHot
	}

	return piefed.CommentSortType(in)
}
