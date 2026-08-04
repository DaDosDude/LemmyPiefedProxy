package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertCommunityBlockView(in piefed.CommunityBlockView) lemmy.CommunityBlockView {
	return lemmy.CommunityBlockView{
		Community: ConvertCommunity(in.Community),
		Person:    ConvertPerson(in.Person),
	}
}
