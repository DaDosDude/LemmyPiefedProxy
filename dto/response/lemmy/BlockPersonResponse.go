package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type BlockPersonResponse struct {
	Blocked    bool             `json:"blocked" required:"true"`
	PersonView lemmy.PersonView `json:"person_view" required:"true"`
}
