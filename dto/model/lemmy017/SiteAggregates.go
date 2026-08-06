package lemmy017

type SiteAggregates struct {
	Id                  uint `json:"id" validate:"required"`
	SiteId              uint `json:"site_id" validate:"required"`
	Users               uint `json:"users" validate:"required"`
	Posts               uint `json:"posts" validate:"required"`
	Comments            uint `json:"comments" validate:"required"`
	Communities         uint `json:"communities" validate:"required"`
	UsersActiveDay      uint `json:"users_active_day" validate:"required"`
	UsersActiveWeek     uint `json:"users_active_week" validate:"required"`
	UsersActiveMonth    uint `json:"users_active_month" validate:"required"`
	UsersActiveHalfYear uint `json:"users_active_half_year" validate:"required"`
}
