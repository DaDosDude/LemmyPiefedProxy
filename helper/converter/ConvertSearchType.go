package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertSearchType(in piefed.SearchType) lemmy.SearchType {
	return lemmy.SearchType(in)
}

// ReverseConvertSearchType: PieFed has no "All" content type. Lemmy's
// SearchTypeAll is mapped to Posts as the most useful single-type default —
// a Lemmy client requesting "All" won't get communities/users/comments back,
// only posts, since PieFed's /search endpoint requires picking one type_.
func ReverseConvertSearchType(in lemmy.SearchType) piefed.SearchType {
	if in == lemmy.SearchTypeAll {
		return piefed.SearchTypePosts
	}

	return piefed.SearchType(in)
}
