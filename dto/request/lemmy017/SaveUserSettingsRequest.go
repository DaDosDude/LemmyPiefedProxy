package lemmy017

// SaveUserSettingsRequest matches Lemmy 0.17.2's real SaveUserSettings —
// confirmed against Lemmy's own source, and genuinely simpler than the
// current-generation client's field set (no default_comment_sort_type,
// no auto_expand/blur_nsfw/collapse_bot_comments/etc. at all — those are
// later additions). DefaultSortType/DefaultListingType are raw numeric
// indices here, unlike every other sort-bearing endpoint.
type SaveUserSettingsRequest struct {
	ShowNsfw                 *bool   `json:"show_nsfw,omitempty"`
	ShowScores               *bool   `json:"show_scores,omitempty"`
	Theme                    *string `json:"theme,omitempty"`
	DefaultSortType          *int16  `json:"default_sort_type,omitempty"`
	DefaultListingType       *int16  `json:"default_listing_type,omitempty"`
	InterfaceLanguage        *string `json:"interface_language,omitempty"`
	Avatar                   *string `json:"avatar,omitempty"`
	Banner                   *string `json:"banner,omitempty"`
	DisplayName              *string `json:"display_name,omitempty"`
	Email                    *string `json:"email,omitempty"`
	Bio                      *string `json:"bio,omitempty"`
	MatrixUserId             *string `json:"matrix_user_id,omitempty"`
	ShowAvatars              *bool   `json:"show_avatars,omitempty"`
	SendNotificationsToEmail *bool   `json:"send_notifications_to_email,omitempty"`
	BotAccount               *bool   `json:"bot_account,omitempty"`
	ShowBotAccounts          *bool   `json:"show_bot_accounts,omitempty"`
	ShowReadPosts            *bool   `json:"show_read_posts,omitempty"`
	ShowNewPostNotifs        *bool   `json:"show_new_post_notifs,omitempty"`
	DiscussionLanguages      []int64 `json:"discussion_languages,omitempty"`
}
