package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertPersonView(in piefed.PersonView) lemmy.PersonView {
	return lemmy.PersonView{
		Counts:  ConvertPersonAggregates(in.Counts),
		IsAdmin: in.IsAdmin,
		Person:  ConvertPerson(in.Person),
	}
}
