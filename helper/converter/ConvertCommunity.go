package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunity(in piefed.Community) lemmy.Community {
	return lemmy.Community{
		ActorId:                 in.ActorId,
		Banner:                  in.Banner,
		Deleted:                 in.Deleted,
		Description:             in.Description,
		Hidden:                  in.Hidden,
		Icon:                    in.Icon,
		Id:                      in.Id,
		InstanceId:              in.InstanceId,
		Local:                   in.Local,
		Name:                    in.Name,
		Nsfw:                    in.Nsfw,
		PostingRestrictedToMods: in.RestrictedToMods,
		Published:               in.Published,
		Removed:                 in.Removed,
		Title:                   in.Title,
		Updated:                 in.Updated,
		Visibility:              lemmy.CommunityVisibilityPublic,
	}
}
