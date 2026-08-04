package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertInstance(in piefed.Instance) lemmy.Instance {
	return lemmy.Instance{
		Domain:    in.Domain,
		Id:        in.Id,
		Published: in.Published,
		Software:  in.Software, // todo simulate
		Updated:   in.Updated,
		Version:   in.Version, // todo simulate
	}
}
