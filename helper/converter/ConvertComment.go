package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertComment(in piefed.Comment) lemmy.Comment {
	return lemmy.Comment{
		ActivityPubId: in.ActivityPubId,
		Content:       in.Body,
		CreatorId:     in.CreatorId,
		Deleted:       in.Deleted,
		Distinguished: in.Distinguished,
		Id:            in.Id,
		LanguageId:    in.LanguageId,
		Local:         in.Local,
		Path:          in.Path,
		PostId:        in.PostId,
		Published:     in.Published,
		Removed:       in.Removed,
		Updated:       in.Updated,
	}
}
