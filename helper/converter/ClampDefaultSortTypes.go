package converter

// Piefed's /user/save_user_settings validates default_sort_type against its
// own default_sorts_list, which is narrower than both the post-listing sort
// set and the general SortType enum — confirmed directly from Piefed's
// schema.go. mlmym's settings dropdown offers many more options (New,
// Controversial, MostComments, TopWeek, etc.) that would 400 here if
// forwarded as-is.
var piefedDefaultSortValues = map[string]bool{
	"Hot": true, "Top": true, "New": true, "Active": true, "Old": true, "Scaled": true,
}

// Similarly narrower for comment sort — confirmed from default_comment_sorts_list.
var piefedDefaultCommentSortValues = map[string]bool{
	"Hot": true, "Top": true, "New": true, "Old": true,
}

func ClampDefaultSortType(in string) string {
	if piefedDefaultSortValues[in] {
		return in
	}
	return "Hot"
}

func ClampDefaultCommentSortType(in string) string {
	if piefedDefaultCommentSortValues[in] {
		return in
	}
	return "Hot"
}
