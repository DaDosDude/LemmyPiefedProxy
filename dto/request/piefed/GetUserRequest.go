package piefed

import "LemmyPiefedApi/dto/model/piefed"

type GetUserRequest struct {
	PersonId  *uint            `json:"person_id,omitempty"`
	Username  *string          `json:"username,omitempty"`
	Sort      *piefed.SortType `json:"sort,omitempty"`
	Page      *uint            `json:"page,omitempty"`
	Limit     *uint            `json:"limit,omitempty"`
	SavedOnly *bool            `json:"saved_only,omitempty"`
}
