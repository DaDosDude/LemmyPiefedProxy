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

	return &lemmy.Site{
		ActorId:         in.ActorId,
		Banner:          nil,
		ContentWarning:  nil,
		Description:     in.Description,
		Icon:            in.Icon,
		Id:              0, // todo
		InboxUrl:        actor.Inbox,
		InstanceId:      0,  // todo
		LastRefreshedAt: "", // todo
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
