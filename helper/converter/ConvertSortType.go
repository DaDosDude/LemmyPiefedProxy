package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertSortType(in piefed.SortType) lemmy.SortType {
	return lemmy.SortType(in)
}

// ReverseConvertSortType: Piefed's own /post/list validation error (seen
// directly from a live request) confirms its accepted sort values are:
// Hot, Top, TopHour, TopSixHour, TopTwelveHour, TopWeek, TopDay, TopMonth,
// TopThreeMonths, TopSixMonths, TopNineMonths, TopYear, TopAll, New, Old,
// Scaled, Active — i.e. Old, TopYear, TopAll, and the TopXMonths variants
// are all directly valid on Piefed and should pass through unchanged, not
// be remapped to something coarser. Only MostComments, Controversial, and
// NewComments are genuinely Lemmy-only values with no Piefed equivalent
// and need translating.
//
// The previous version of this function used a Go switch statement
// written as if cases fall through by default (as in C/JavaScript) — they
// don't in Go. Each bare `case X:` with no body does nothing and exits the
// switch, so SortTypeOld and SortTypeMostComments were silently falling
// through to the plain passthrough at the bottom instead of being grouped
// with SortTypeControversial as originally intended. Old passing through
// happened to still work (it's a valid Piefed value), but MostComments
// isn't, and calling with it produced a hard 400 on every request.
func ReverseConvertSortType(in lemmy.SortType) piefed.SortType {
	switch in {
	case lemmy.SortTypeMostComments, lemmy.SortTypeControversial:
		return piefed.SortTypeActive
	case lemmy.SortTypeNewComments:
		return piefed.SortTypeNew
	}

	return piefed.SortType(in)
}
