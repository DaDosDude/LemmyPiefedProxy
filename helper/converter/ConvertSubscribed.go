package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertSubscribedType(in piefed.SubscribedType) lemmy.SubscribedType {
	return lemmy.SubscribedType(in)
}
