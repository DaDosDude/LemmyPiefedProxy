package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunityFollowerView(in piefed.CommunityFollowerView) lemmy.CommunityFollowerView {
	return lemmy.CommunityFollowerView{
		Community: ConvertCommunity(in.Community),
		Follower:  ConvertPerson(in.Follower),
	}
}
