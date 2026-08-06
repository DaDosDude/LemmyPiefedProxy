package lemmy017

type LocalUserSettingsView struct {
	LocalUser LocalUserSettings `json:"local_user" validate:"required"`
	Person    Person            `json:"person" validate:"required"`
	Counts    PersonAggregates  `json:"counts" validate:"required"`
}
