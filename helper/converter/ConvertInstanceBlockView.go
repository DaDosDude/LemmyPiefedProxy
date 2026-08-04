package converter

import (
	"LemmyBeProxy/dto/model/ap"
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertInstanceBlockView(in piefed.InstanceBlockView, siteActor ap.Actor) lemmy.InstanceBlockView {
	return lemmy.InstanceBlockView{
		Instance: ConvertInstance(in.Instance),
		Person:   ConvertPerson(in.Person),
		Site:     ConvertSite(in.Site, siteActor),
	}
}
