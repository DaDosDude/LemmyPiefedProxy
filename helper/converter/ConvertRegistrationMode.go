package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertRegistrationMode(in piefed.RegistrationMode) lemmy.RegistrationMode {
	return lemmy.RegistrationMode(in)
}
