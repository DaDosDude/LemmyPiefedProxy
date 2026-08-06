package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// SiteView matches Lemmy 0.17.2's real SiteView. Site is reused directly
// from canonical — confirmed field-by-field clean, unlike LocalSite,
// SiteAggregates, and LocalSiteRateLimit, which all needed real changes.
type SiteView struct {
	Site               lemmy.Site         `json:"site" validate:"required"`
	LocalSite          LocalSite          `json:"local_site" validate:"required"`
	LocalSiteRateLimit LocalSiteRateLimit `json:"local_site_rate_limit" validate:"required"`
	Counts             SiteAggregates     `json:"counts" validate:"required"`
}
