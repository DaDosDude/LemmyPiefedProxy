package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertPost(in piefed.Post) lemmy.Post {
	return lemmy.Post{
		AltText:           in.AltText,
		ActivityPubId:     in.ActivityPubId,
		Body:              in.Body,
		CommunityId:       in.CommunityId,
		CreatorId:         in.CreatorId,
		Deleted:           in.Deleted,
		EmbedDescription:  nil,
		EmbedTitle:        nil,
		EmbedVideoUrl:     nil,
		FeaturedCommunity: false,
		FeaturedLocal:     false,
		Id:                in.Id,
		LanguageId:        in.LanguageId,
		Local:             in.Local,
		Locked:            in.Locked,
		Name:              in.Title,
		Nsfw:              in.Nsfw,
		Published:         in.Published,
		Removed:           in.Removed,
		ThumbnailUrl:      in.ThumbnailUrl,
		Updated:           in.Updated,
		Url:               in.Url,
		UrlContentType:    nil,
	}
}
