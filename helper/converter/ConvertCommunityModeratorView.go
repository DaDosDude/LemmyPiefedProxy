package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunityModeratorView(in piefed.CommunityModeratorView) lemmy.CommunityModeratorView {
	return lemmy.CommunityModeratorView{
		Community: ConvertCommunity(in.Community),
		Moderator: ConvertPerson(in.Moderator),
	}
}
