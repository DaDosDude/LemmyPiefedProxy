package lemmy017

// PostAggregates matches Lemmy 0.17.2's real PostAggregates struct.
// Confirmed missing from our canonical PostAggregates: Id (the
// aggregates row's own primary key, distinct from PostId — defaulted to
// PostId's value, a harmless stand-in with no real semantic meaning to
// API consumers beyond existing), NewestCommentTimeNecro (a
// necro-bump-prevention variant of NewestCommentTime capped to 2 days —
// no equivalent tracked at all, approximated by reusing
// NewestCommentTime, an imperfect but harmless stand-in), and
// FeaturedCommunity/FeaturedLocal (real data, just tracked on the
// sibling Post object in our canonical model rather than here).
type PostAggregates struct {
	Id                      uint   `json:"id" validate:"required"`
	PostId                  uint   `json:"post_id" validate:"required"`
	Comments                uint   `json:"comments" validate:"required"`
	Score                   int    `json:"score" validate:"required"`
	Upvotes                 int    `json:"upvotes" validate:"required"`
	Downvotes               int    `json:"downvotes" validate:"required"`
	Published               string `json:"published" validate:"required"`
	NewestCommentTimeNecro  string `json:"newest_comment_time_necro" validate:"required"`
	NewestCommentTime       string `json:"newest_comment_time" validate:"required"`
	FeaturedCommunity       bool   `json:"featured_community" validate:"required"`
	FeaturedLocal           bool   `json:"featured_local" validate:"required"`
}
