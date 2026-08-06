package lemmy017

// LocalUserSettings matches Lemmy 0.17.2's real LocalUserSettings.
// DefaultSortType/DefaultListingType are raw i16 enum indices in real
// 0.17.x — see converter.ConvertSortTypeToIndex017 for the mapping,
// including how newer sort/listing values with no 0.17.2 equivalent are
// handled. ValidatorTime and ShowNewPostNotifs are fields our canonical
// model never tracked at all — defaulted (see the converter), an honest
// gap rather than a guess.
type LocalUserSettings struct {
	Id                       uint    `json:"id" validate:"required"`
	PersonId                 uint    `json:"person_id" validate:"required"`
	Email                    *string `json:"email,omitempty"`
	ShowNsfw                 bool    `json:"show_nsfw" validate:"required"`
	Theme                    string  `json:"theme" validate:"required"`
	DefaultSortType          int16   `json:"default_sort_type" validate:"required"`
	DefaultListingType       int16   `json:"default_listing_type" validate:"required"`
	InterfaceLanguage        string  `json:"interface_language" validate:"required"`
	ShowAvatars              bool    `json:"show_avatars" validate:"required"`
	SendNotificationsToEmail bool    `json:"send_notifications_to_email" validate:"required"`
	ValidatorTime            string  `json:"validator_time" validate:"required"`
	ShowBotAccounts          bool    `json:"show_bot_accounts" validate:"required"`
	ShowScores               bool    `json:"show_scores" validate:"required"`
	ShowReadPosts            bool    `json:"show_read_posts" validate:"required"`
	ShowNewPostNotifs        bool    `json:"show_new_post_notifs" validate:"required"`
	EmailVerified            bool    `json:"email_verified" validate:"required"`
	AcceptedApplication      bool    `json:"accepted_application" validate:"required"`
}
