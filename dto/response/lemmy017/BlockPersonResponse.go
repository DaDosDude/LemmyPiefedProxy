package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type BlockPersonResponse struct {
	Blocked    bool                `json:"blocked" validate:"required"`
	PersonView model017.PersonView `json:"person_view" validate:"required"`
}
