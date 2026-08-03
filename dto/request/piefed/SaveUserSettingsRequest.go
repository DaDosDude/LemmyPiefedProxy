package piefed

type SaveUserSettingsRequest struct {
	ShowNsfw               *bool   `json:"show_nsfw,omitempty"`
	DefaultSortType        *string `json:"default_sort_type,omitempty"`
	DefaultCommentSortType *string `json:"default_comment_sort_type,omitempty"`
	ShowReadPosts          *bool   `json:"show_read_posts,omitempty"`
}
