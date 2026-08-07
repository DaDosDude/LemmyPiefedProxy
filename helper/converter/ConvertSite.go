package converter

import (
	"LemmyBeProxy/dto/model/ap"
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertSite(in *piefed.Site, actor ap.Actor) *lemmy.Site {
	if in == nil {
		return nil
	}

	// LastRefreshedAt was a genuine, marked "todo" left as an empty
	// string until now — real Lemmy's Site struct requires this to be a
	// valid datetime, and an empty string fails to deserialize entirely
	// on strict Rust/serde clients (confirmed directly: this broke
	// lemmyBB with a parse error, since Rust's chrono can't parse an
	// empty string as a datetime). actor.Updated is the closest real
	// equivalent available — it's literally when this ActivityPub
	// actor's data was last updated, which is what this field means.
	// It's optional (not every actor has been updated since creation),
	// so fall back to Published, which every actor has.
	lastRefreshedAt := actor.Published
	if actor.Updated != nil {
		lastRefreshedAt = *actor.Updated
	}

	return &lemmy.Site{
		ActorId:         in.ActorId,
		Banner:          nil,
		ContentWarning:  nil,
		Description:     in.Description,
		Icon:            in.Icon,
		Id:              0, // todo
		InboxUrl:        actor.Inbox,
		InstanceId:      0,  // todo
		LastRefreshedAt: lastRefreshedAt,
		Name:            in.Name,
		PublicKey:       actor.PublicKey.PublicKeyPem,
		Published:       actor.Published,
		Sidebar:         in.Sidebar,
		Updated:         actor.Updated,
	}
}

func ConvertSiteToView(in *piefed.Site, actor ap.Actor) lemmy.SiteView {
	return lemmy.SiteView{
		LocalSite: lemmy.LocalSite{
			ActorNameMaxLength:         20,
			ApplicationEmailAdmins:     false,
			ApplicationQuestion:        nil,
			CaptchaDifficulty:          "",
			CaptchaEnabled:             false,
			CommunityCreationAdminOnly: false,
			DefaultPostListingMode:     lemmy.PostListingModeList,
			DefaultPostListingType:     lemmy.ListingTypeSubscribed,
			DefaultSortType:            lemmy.SortTypeHot,
			DefaultTheme:               "browser",
			EnableDownvotes:            *in.EnableDownvotes,
			EnableNsfw:                 true,
			FederationEnabled:          true,
			FederationSignedFetch:      false,
			HideModlogModNames:         false,
			Id:                         0, // todo
			LegalInformation:           nil,
			PrivateInstance:            false,
			Published:                  actor.Published,
			RegistrationMode:           ConvertRegistrationMode(*in.RegistrationMode),
			ReportsEmailAdmin:          false,
			ReportsEmailVerification:   false,
			SiteId:                     0, // todo
			SiteSetup:                  true,
			SlurFilterRegex:            nil,
			Updated:                    actor.Updated,
		},
		LocalSiteRateLimit: lemmy.LocalSiteRateLimit{
			Published: actor.Published,
		},
		Site: *ConvertSite(in, actor),
	}
}
