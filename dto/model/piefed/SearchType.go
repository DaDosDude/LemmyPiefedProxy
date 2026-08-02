package piefed

// PieFed's content_type_list has no "All" option (unlike Lemmy's SearchType) —
// only Communities, Posts, Users, Url, Comments are valid.
type SearchType string

const (
	SearchTypeComments    SearchType = "Comments"
	SearchTypePosts       SearchType = "Posts"
	SearchTypeCommunities SearchType = "Communities"
	SearchTypeUsers       SearchType = "Users"
	SearchTypeUrl         SearchType = "Url"
)
