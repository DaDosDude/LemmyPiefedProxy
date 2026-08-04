package piefed

import "LemmyBeProxy/dto/model/piefed"

type BlockPersonResponse struct {
	Blocked    bool               `json:"blocked" required:"true"`
	PersonView piefed.PersonView `json:"person_view" required:"true"`
}
