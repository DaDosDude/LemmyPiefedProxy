package lemmy

type SearchType string

const (
	SearchTypeAll         SearchType = "All"
	SearchTypeComments    SearchType = "Comments"
	SearchTypePosts       SearchType = "Posts"
	SearchTypeCommunities SearchType = "Communities"
	SearchTypeUsers       SearchType = "Users"
	SearchTypeUrl         SearchType = "Url"
)
