package lemmy

// Only ShowNsfw, DefaultSortType, DefaultCommentSortType, and ShowReadPosts
// have a real Piefed equivalent (confirmed against Piefed's own
// UserSaveSettingsRequest schema, which itself notes "not all settings
// implemented yet"). Every other field here is accepted — so mlmym's save
// request doesn't fail outright — but has nothing to forward to and is
// silently dropped. InfiniteScrollEnabled in particular can never persist
// server-side on a Piefed-backed instance: Piefed has no field for it at
// all, not a translation gap, a genuine absence on Piefed's side.
type SaveUserSettingsRequest struct {
	ShowNsfw                 *bool    `json:"show_nsfw,omitempty"`
	BlurNsfw                 *bool    `json:"blur_nsfw,omitempty"`
	AutoExpand               *bool    `json:"auto_expand,omitempty"`
	Theme                    *string  `json:"theme,omitempty"`
	DefaultSortType          *string  `json:"default_sort_type,omitempty"`
	DefaultListingType       *string  `json:"default_listing_type,omitempty"`
	InterfaceLanguage        *string  `json:"interface_language,omitempty"`
	EnableKeyboardNavigation *bool    `json:"enable_keyboard_navigation,omitempty"`
	InfiniteScrollEnabled    *bool    `json:"infinite_scroll_enabled,omitempty"`
	ShowAvatars              *bool    `json:"show_avatars,omitempty"`
	EnableAnimatedImages     *bool    `json:"enable_animated_images,omitempty"`
	ShowScores               *bool    `json:"show_scores,omitempty"`
	ShowUpvotes              *bool    `json:"show_upvotes,omitempty"`
	ShowDownvotes            *bool    `json:"show_downvotes,omitempty"`
	ShowUpvotePercentage     *bool    `json:"show_upvote_percentage,omitempty"`
	SendNotificationsToEmail *bool    `json:"send_notifications_to_email,omitempty"`
	ShowBotAccounts          *bool    `json:"show_bot_accounts,omitempty"`
	CollapseBotComments      *bool    `json:"collapse_bot_comments,omitempty"`
	ShowReadPosts            *bool    `json:"show_read_posts,omitempty"`
	DiscussionLanguages      []int64  `json:"discussion_languages,omitempty"`
	OpenLinksInNewTab        *bool    `json:"open_links_in_new_tab,omitempty"`
	DefaultCommentSortType   *string  `json:"default_comment_sort_type,omitempty"`
}
