package piefed

// ResolveObjectRequest — Piefed's own release notes confirm they
// implemented resolve_object, but their exact route path and field
// naming weren't directly verifiable (codeberg.org blocks automated
// fetches). This assumes /api/alpha/resolve_object?q=... matching
// every other Piefed endpoint's established 1:1 mirroring of Lemmy's
// own shape and naming convention — worth confirming directly against
// a live instance once deployed.
type ResolveObjectRequest struct {
	Q string `json:"q"`
}
