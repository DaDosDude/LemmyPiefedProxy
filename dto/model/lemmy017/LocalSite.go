package lemmy017

// LocalSite matches Lemmy 0.17.2's real LocalSite. Three fields our
// canonical model never tracked at all — RequireEmailVerification,
// FederationDebug, FederationWorkerCount — default to false/0, an
// honest gap rather than a guess. DefaultPostListingMode,
// FederationSignedFetch, and ReportsEmailVerification are later
// additions canonical has that 0.17.2 doesn't, so they're simply not
// present here.
type LocalSite struct {
	Id                         uint    `json:"id" validate:"required"`
	SiteId                     uint    `json:"site_id" validate:"required"`
	SiteSetup                  bool    `json:"site_setup" validate:"required"`
	EnableDownvotes            bool    `json:"enable_downvotes" validate:"required"`
	EnableNsfw                 bool    `json:"enable_nsfw" validate:"required"`
	CommunityCreationAdminOnly bool    `json:"community_creation_admin_only" validate:"required"`
	RequireEmailVerification   bool    `json:"require_email_verification" validate:"required"`
	ApplicationQuestion        *string `json:"application_question"`
	PrivateInstance            bool    `json:"private_instance" validate:"required"`
	DefaultTheme               string  `json:"default_theme" validate:"required"`
	DefaultPostListingType     string  `json:"default_post_listing_type" validate:"required"`
	LegalInformation           *string `json:"legal_information"`
	HideModlogModNames         bool    `json:"hide_modlog_mod_names" validate:"required"`
	ApplicationEmailAdmins     bool    `json:"application_email_admins" validate:"required"`
	SlurFilterRegex            *string `json:"slur_filter_regex"`
	ActorNameMaxLength         uint    `json:"actor_name_max_length" validate:"required"`
	FederationEnabled          bool    `json:"federation_enabled" validate:"required"`
	FederationDebug            bool    `json:"federation_debug" validate:"required"`
	FederationWorkerCount      uint    `json:"federation_worker_count" validate:"required"`
	CaptchaEnabled             bool    `json:"captcha_enabled" validate:"required"`
	CaptchaDifficulty          string  `json:"captcha_difficulty" validate:"required"`
	RegistrationMode           string  `json:"registration_mode" validate:"required"`
	ReportsEmailAdmins         bool    `json:"reports_email_admins" validate:"required"`
	Published                  string  `json:"published" validate:"required"`
	Updated                    *string `json:"updated"`
}
