package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertPersonBlockView(in piefed.PersonBlockView) lemmy.PersonBlockView {
	return lemmy.PersonBlockView{
		Person: ConvertPerson(in.Person),
		Target: ConvertPerson(in.Target),
	}
}
