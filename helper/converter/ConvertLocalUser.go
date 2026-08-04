package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
	"time"
)

func ConvertLocalUser(in piefed.LocalUser, person piefed.Person) lemmy.LocalUser {
	return lemmy.LocalUser{
		AcceptedApplication:      true,
		Admin:                    false, // todo
		AutoExpand:               false,
		BlurNsfw:                 false,
		CollapseBotComments:      false,
		DefaultListingType:       ConvertListingType(in.DefaultListingType),
		DefaultSortType:          ConvertSortType(in.DefaultSortType),
		Email:                    nil,
		EmailVerified:            true,
		EnableAnimatedImages:     true,
		EnableKeyboardNavigation: true,
		Id:                       0, // todo
		InfiniteScrollEnabled:    true,
		InterfaceLanguage:        "en", // todo
		LastDonationNotification: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		OpenLinksInNewTab:        true,
		PersonId:                 person.Id,
		PostListingMode:          lemmy.PostListingModeList,
		SendNotificationsToEmail: false,
		ShowAvatars:              true,
		ShowBotAccounts:          in.ShowBotAccounts,
		ShowNsfw:                 in.ShowNsfw,
		ShowReadPosts:            in.ShowReadPosts,
		ShowScores:               in.ShowScores,
		Theme:                    "browser",
		TotpEnabled:              false,
	}
}
