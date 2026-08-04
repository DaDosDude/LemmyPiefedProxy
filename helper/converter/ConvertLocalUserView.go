package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertLocalUserView(in piefed.LocalUserView) lemmy.LocalUserView {
	return lemmy.LocalUserView{
		Counts:    ConvertPersonAggregates(in.Counts),
		LocalUser: ConvertLocalUser(in.LocalUser, in.Person),
		LocalUserVoteDisplayMode: lemmy.LocalUserVoteDisplayMode{
			Downvotes:        true,
			LocalUserId:      0,
			Score:            true,
			UpvotePercentage: true,
			Upvotes:          true,
		},
		Person: ConvertPerson(in.Person),
	}
}
