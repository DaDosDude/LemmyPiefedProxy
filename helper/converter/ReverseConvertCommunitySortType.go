package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

// piefedCommunitySortValues is Piefed's actual accepted set of sort values
// for /community/list — confirmed directly from Piefed's own validation
// error ("Must be one of: Hot, Top, New, Old, Active, TopAll, TopPosts,
// TopSubscribers, NewFederated, OldFederated"), not from its schema.py
// source, which listed a different (and apparently not what's actually
// enforced at runtime) set.
//
// Crucially, this set has no per-hour/day/week/month Top* granularity and
// no Scaled — both valid for post listing. A Lemmy client's sort
// preference (e.g. "Scaled", commonly used for posts) forwarded unchanged
// to community listing gets a hard 400 on every single request, since
// community listing has no such option at all.
var piefedCommunitySortValues = map[string]bool{
	"Hot": true, "Top": true, "New": true, "Old": true, "Active": true,
	"TopAll": true, "TopPosts": true, "TopSubscribers": true,
	"NewFederated": true, "OldFederated": true,
}

// ReverseConvertCommunitySortType is the community-listing-specific
// counterpart to ReverseConvertSortType. Any Lemmy sort value Piefed's
// community endpoint doesn't actually accept falls back to Hot rather than
// being forwarded verbatim and rejected.
func ReverseConvertCommunitySortType(in lemmy.SortType) piefed.SortType {
	if piefedCommunitySortValues[string(in)] {
		return piefed.SortType(in)
	}
	return piefed.SortTypeHot
}
