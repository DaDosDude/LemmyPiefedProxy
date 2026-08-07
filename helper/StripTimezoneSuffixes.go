package helper

import "regexp"

// timestampWithTimezone matches an RFC3339-shaped timestamp with a
// trailing timezone offset ("+00:00") or "Z" — the format this proxy's
// canonical model uses throughout. Capture group 1 is the "naive"
// portion (no timezone) that real Lemmy 0.17.x actually expects.
var timestampWithTimezone = regexp.MustCompile(
	`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?)(Z|[+-]\d{2}:\d{2})`,
)

// StripTimezoneSuffixes removes the timezone portion from every
// RFC3339-shaped timestamp in a JSON response body.
//
// This exists specifically for FRONTEND_VERSION=0.17: real Lemmy 0.17.x
// uses chrono::NaiveDateTime for every one of its timestamp fields
// (published, updated, last_refreshed_at, validator_time — confirmed
// systemic across comment, community, post, person, site, and every
// aggregates type in Lemmy's own 0.17.2 source), and NaiveDateTime has
// no timezone concept at all. Its deserializer doesn't just ignore an
// unexpected timezone suffix — it errors on the leftover characters,
// which serde_json surfaces as a "trailing input" parse failure. A real
// lemmyBB client hits this on essentially every response with a
// timestamp, which is most of them.
//
// This runs as a single regex pass over the fully marshaled JSON body
// rather than being threaded through every individual 0.17-shaped type
// and converter function — the pattern is specific enough (a full
// RFC3339 date-time, not just any string) that it won't misfire on
// unrelated content, and one central pass is far less error-prone than
// remembering to apply it at dozens of call sites across every response
// type this proxy builds.
func StripTimezoneSuffixes(jsonBody []byte) []byte {
	return timestampWithTimezone.ReplaceAll(jsonBody, []byte("$1"))
}
