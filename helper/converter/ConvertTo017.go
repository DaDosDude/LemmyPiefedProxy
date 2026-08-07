package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/lemmy017"
	"LemmyBeProxy/helper"
	"strings"
)

// ConvertPersonTo017 fills in the two fields 0.17.2's real PersonSafe
// requires that our canonical Person doesn't track — see
// dto/model/lemmy017/Person.go for why each is handled the way it is.
func ConvertPersonTo017(in lemmy.Person) lemmy017.Person {
	botAccount := false
	if in.BotAccount != nil {
		botAccount = *in.BotAccount
	}

	return lemmy017.Person{
		Id:           in.Id,
		Name:         in.Name,
		DisplayName:  strPtrOrNil(in.DisplayName),
		Avatar:       in.Avatar,
		Banned:       in.Banned,
		Published:    in.Published,
		Updated:      in.Updated,
		ActorId:      in.ActorId,
		Bio:          in.Bio,
		Local:        in.Local,
		Banner:       in.Banner,
		Deleted:      in.Deleted,
		InboxUrl:     in.ActorId + "/inbox",
		MatrixUserId: in.MatrixUserId,
		Admin:        false,
		BotAccount:   botAccount,
		InstanceId:   in.InstanceId,
	}
}

// strPtrOrNil converts Person.DisplayName (a plain string in the
// canonical model, empty string meaning "not set") into 0.17.x's
// Option<String> convention (nil meaning "not set").
func strPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ConvertPostAggregatesTo017 fills in the fields 0.17.2's real
// PostAggregates requires that our canonical model doesn't track — see
// dto/model/lemmy017/PostAggregates.go for why each is handled the way
// it is. featuredCommunity/featuredLocal come from the sibling Post
// object, since that's real data, just tracked in a different place in
// the canonical model.
func ConvertPostAggregatesTo017(in lemmy.PostAggregates, featuredCommunity bool, featuredLocal bool) lemmy017.PostAggregates {
	return lemmy017.PostAggregates{
		Id:                     in.PostId,
		PostId:                 in.PostId,
		Comments:               in.Comments,
		Score:                  in.Score,
		Upvotes:                in.Upvotes,
		Downvotes:              in.Downvotes,
		Published:              in.Published,
		NewestCommentTimeNecro: in.NewestCommentTime,
		NewestCommentTime:      in.NewestCommentTime,
		FeaturedCommunity:      featuredCommunity,
		FeaturedLocal:          featuredLocal,
	}
}

// ConvertCommentAggregatesTo017 fills in the field 0.17.2's real
// CommentAggregates requires that our canonical model doesn't track —
// see dto/model/lemmy017/CommentAggregates.go for why.
func ConvertCommentAggregatesTo017(in lemmy.CommentAggregates) lemmy017.CommentAggregates {
	return lemmy017.CommentAggregates{
		Id:         in.CommentId,
		CommentId:  in.CommentId,
		Score:      in.Score,
		Upvotes:    in.Upvotes,
		Downvotes:  in.Downvotes,
		Published:  in.Published,
		ChildCount: in.ChildCount,
	}
}

// ConvertCommentViewTo017 assembles a full 0.17.x-shaped CommentView from
// the canonical one.
func ConvertCommentViewTo017(in lemmy.CommentView) lemmy017.CommentView {
	return lemmy017.CommentView{
		Comment:                    in.Comment,
		Creator:                    ConvertPersonTo017(in.Creator),
		Post:                       in.Post,
		Community:                  in.Community,
		Counts:                     ConvertCommentAggregatesTo017(in.Counts),
		CreatorBannedFromCommunity: in.CreatorBannedFromCommunity,
		Subscribed:                 in.Subscribed,
		Saved:                      in.Saved,
		CreatorBlocked:             in.CreatorBlocked,
		MyVote:                     in.MyVote,
	}
}

// ConvertCommunityAggregatesTo017 fills in the field 0.17.2's real
// CommunityAggregates requires that our canonical model doesn't track —
// same extra-id pattern as Post/CommentAggregates.
func ConvertCommunityAggregatesTo017(in lemmy.CommunityAggregates) lemmy017.CommunityAggregates {
	return lemmy017.CommunityAggregates{
		Id:                  in.CommunityId,
		CommunityId:         in.CommunityId,
		Subscribers:         in.Subscribers,
		Posts:               in.Posts,
		Comments:            in.Comments,
		Published:           in.Published,
		UsersActiveDay:      in.UsersActiveDay,
		UsersActiveWeek:     in.UsersActiveWeek,
		UsersActiveMonth:    in.UsersActiveMonth,
		UsersActiveHalfYear: in.UsersActiveHalfYear,
	}
}

// ConvertCommunityViewTo017 assembles a full 0.17.x-shaped CommunityView
// from the canonical one.
func ConvertCommunityViewTo017(in lemmy.CommunityView) lemmy017.CommunityView {
	return lemmy017.CommunityView{
		Community:  in.Community,
		Subscribed: in.Subscribed,
		Blocked:    in.Blocked,
		Counts:     ConvertCommunityAggregatesTo017(in.Counts),
	}
}

// ConvertPersonAggregatesTo017 fills in the fields 0.17.2's real
// PersonAggregates requires that our canonical model doesn't track at
// all — Id (same extra-id pattern as elsewhere), PostScore, and
// CommentScore (fields never captured, not just omitted; both default
// to 0, an honest gap rather than a guess).
func ConvertPersonAggregatesTo017(in lemmy.PersonAggregates) lemmy017.PersonAggregates {
	return lemmy017.PersonAggregates{
		Id:           in.PersonId,
		PersonId:     in.PersonId,
		PostCount:    in.PostCount,
		PostScore:    0,
		CommentCount: in.CommentCount,
		CommentScore: 0,
	}
}

// ConvertPersonViewTo017 assembles a full 0.17.x-shaped PersonView from
// the canonical one.
func ConvertPersonViewTo017(in lemmy.PersonView) lemmy017.PersonView {
	return lemmy017.PersonView{
		Person: ConvertPersonTo017(in.Person),
		Counts: ConvertPersonAggregatesTo017(in.Counts),
	}
}

// ConvertSiteAggregatesTo017 fills in the extra id field 0.17.2's real
// SiteAggregates requires, same pattern as other Aggregates types.
func ConvertSiteAggregatesTo017(in lemmy.SiteAggregates) lemmy017.SiteAggregates {
	return lemmy017.SiteAggregates{
		Id:                  in.SiteId,
		SiteId:              in.SiteId,
		Users:               in.Users,
		Posts:               in.Posts,
		Comments:            in.Comments,
		Communities:         in.Communities,
		UsersActiveDay:      in.UsersActiveDay,
		UsersActiveWeek:     in.UsersActiveWeek,
		UsersActiveMonth:    in.UsersActiveMonth,
		UsersActiveHalfYear: in.UsersActiveHalfYear,
	}
}

// ConvertLocalSiteRateLimitTo017 fills in the extra id field 0.17.2's
// real LocalSiteRateLimit requires. ImportUserSettings fields are
// simply dropped — 0.17.2 has no such concept.
func ConvertLocalSiteRateLimitTo017(in lemmy.LocalSiteRateLimit) lemmy017.LocalSiteRateLimit {
	return lemmy017.LocalSiteRateLimit{
		Id:                in.LocalSiteId,
		LocalSiteId:       in.LocalSiteId,
		Message:           in.Message,
		MessagePerSecond:  in.MessagePerSecond,
		Post:              in.Post,
		PostPerSecond:     in.PostPerSecond,
		Register:          in.Register,
		RegisterPerSecond: in.RegisterPerSecond,
		Image:             in.Image,
		ImagePerSecond:    in.ImagePerSecond,
		Comment:           in.Comment,
		CommentPerSecond:  in.CommentPerSecond,
		Search:            in.Search,
		SearchPerSecond:   in.SearchPerSecond,
		Published:         in.Published,
		Updated:           in.Updated,
	}
}

// ConvertLocalSiteTo017 handles LocalSite's three genuinely missing
// fields (defaulted honestly, see lemmy017.LocalSite's own comment) and
// converts DefaultPostListingType/RegistrationMode from named string
// types to plain strings for the wire.
func ConvertLocalSiteTo017(in lemmy.LocalSite) lemmy017.LocalSite {
	return lemmy017.LocalSite{
		Id:                         in.Id,
		SiteId:                     in.SiteId,
		SiteSetup:                  in.SiteSetup,
		EnableDownvotes:            in.EnableDownvotes,
		EnableNsfw:                 in.EnableNsfw,
		CommunityCreationAdminOnly: in.CommunityCreationAdminOnly,
		RequireEmailVerification:   false,
		ApplicationQuestion:        in.ApplicationQuestion,
		PrivateInstance:            in.PrivateInstance,
		DefaultTheme:               in.DefaultTheme,
		DefaultPostListingType:     string(in.DefaultPostListingType),
		LegalInformation:           in.LegalInformation,
		HideModlogModNames:         in.HideModlogModNames,
		ApplicationEmailAdmins:     in.ApplicationEmailAdmins,
		SlurFilterRegex:            in.SlurFilterRegex,
		ActorNameMaxLength:         in.ActorNameMaxLength,
		FederationEnabled:          in.FederationEnabled,
		FederationDebug:            false,
		FederationWorkerCount:      0,
		CaptchaEnabled:             in.CaptchaEnabled,
		CaptchaDifficulty:          in.CaptchaDifficulty,
		// Confirmed the ONLY enum in real 0.17.2's entire source with
		// #[serde(rename_all = "lowercase")] — every other enum
		// (SortType, ListingType, CommentSortType, SubscribedType, etc.)
		// uses standard PascalCase like canonical does, so this
		// lowercasing is specific to RegistrationMode alone.
		RegistrationMode: strings.ToLower(string(in.RegistrationMode)),
		ReportsEmailAdmins:         in.ReportsEmailAdmin,
		Published:                  in.Published,
		Updated:                    in.Updated,
	}
}

// ConvertSiteViewTo017 assembles a full 0.17.x-shaped SiteView from the
// canonical one.
func ConvertSiteViewTo017(in lemmy.SiteView) lemmy017.SiteView {
	return lemmy017.SiteView{
		Site:               in.Site,
		LocalSite:          ConvertLocalSiteTo017(in.LocalSite),
		LocalSiteRateLimit: ConvertLocalSiteRateLimitTo017(in.LocalSiteRateLimit),
		Counts:             ConvertSiteAggregatesTo017(in.Counts),
	}
}

// sortTypeTo017Index maps canonical SortType to Lemmy 0.17.2's real
// numeric enum ordering (confirmed against Lemmy's own source: Active,
// Hot, New, Old, TopDay, TopWeek, TopMonth, TopYear, TopAll,
// MostComments, NewComments — in that order, 0-indexed). Several
// canonical values have no 0.17.2 equivalent at all (Scaled,
// Controversial, TopHour, TopSixHour, TopTwelveHour, TopThreeMonths,
// TopSixMonths, TopNineMonths — all added in later Lemmy versions) and
// map to the closest reasonable existing value rather than an arbitrary
// default.
var sortTypeTo017Index = map[lemmy.SortType]int16{
	lemmy.SortTypeActive:        0,
	lemmy.SortTypeHot:           1,
	lemmy.SortTypeNew:           2,
	lemmy.SortTypeOld:           3,
	lemmy.SortTypeTopDay:        4,
	lemmy.SortTypeTopWeek:       5,
	lemmy.SortTypeTopMonth:      6,
	lemmy.SortTypeTopYear:       7,
	lemmy.SortTypeTopAll:        8,
	lemmy.SortTypeMostComments:  9,
	lemmy.SortTypeNewComments:   10,
	// No 0.17.2 equivalent — mapped to the closest existing value.
	lemmy.SortTypeScaled:         0, // Active
	lemmy.SortTypeControversial:  1, // Hot
	lemmy.SortTypeTopHour:        4, // TopDay
	lemmy.SortTypeTopSixHour:     4, // TopDay
	lemmy.SortTypeTopTwelveHour:  4, // TopDay
	lemmy.SortTypeTopThreeMonths: 6, // TopMonth
	lemmy.SortTypeTopSixMonths:   6, // TopMonth
	lemmy.SortTypeTopNineMonths:  6, // TopMonth
}

// ConvertSortTypeToIndex017 converts a canonical SortType to 0.17.2's
// real numeric enum index. Falls back to Hot (1) for anything
// unrecognized, same fallback philosophy as ClampDefaultSortType.
func ConvertSortTypeToIndex017(in lemmy.SortType) int16 {
	if index, ok := sortTypeTo017Index[in]; ok {
		return index
	}
	return 1
}

// listingTypeTo017Index maps canonical ListingType to 0.17.2's real
// numeric ordering (All, Local, Subscribed). ModeratorView is a later
// addition with no 0.17.2 equivalent, mapped to All.
var listingTypeTo017Index = map[lemmy.ListingType]int16{
	lemmy.ListingTypeAll:        0,
	lemmy.ListingTypeLocal:      1,
	lemmy.ListingTypeSubscribed: 2,
	lemmy.ListingTypeModeratorView: 0, // All
}

// ConvertListingTypeToIndex017 converts a canonical ListingType to
// 0.17.2's real numeric enum index. Falls back to All (0).
func ConvertListingTypeToIndex017(in lemmy.ListingType) int16 {
	if index, ok := listingTypeTo017Index[in]; ok {
		return index
	}
	return 0
}

func ConvertCommunityModeratorViewTo017(in lemmy.CommunityModeratorView) lemmy017.CommunityModeratorView {
	return lemmy017.CommunityModeratorView{
		Community: in.Community,
		Moderator: ConvertPersonTo017(in.Moderator),
	}
}

func ConvertCommunityFollowerViewTo017(in lemmy.CommunityFollowerView) lemmy017.CommunityFollowerView {
	return lemmy017.CommunityFollowerView{
		Community: in.Community,
		Follower:  ConvertPersonTo017(in.Follower),
	}
}

func ConvertCommunityBlockViewTo017(in lemmy.CommunityBlockView) lemmy017.CommunityBlockView {
	return lemmy017.CommunityBlockView{
		Person:    ConvertPersonTo017(in.Person),
		Community: in.Community,
	}
}

func ConvertPersonBlockViewTo017(in lemmy.PersonBlockView) lemmy017.PersonBlockView {
	return lemmy017.PersonBlockView{
		Person: ConvertPersonTo017(in.Person),
		Target: ConvertPersonTo017(in.Target),
	}
}

// ConvertLocalUserSettingsTo017 handles the numeric sort/listing type
// encoding and the two fields our canonical model never tracked at all
// (ValidatorTime, ShowNewPostNotifs) — see lemmy017.LocalUserSettings's
// own comment.
func ConvertLocalUserSettingsTo017(in lemmy.LocalUser) lemmy017.LocalUserSettings {
	return lemmy017.LocalUserSettings{
		Id:                       in.Id,
		PersonId:                 in.PersonId,
		Email:                    in.Email,
		ShowNsfw:                 in.ShowNsfw,
		Theme:                    in.Theme,
		DefaultSortType:          ConvertSortTypeToIndex017(in.DefaultSortType),
		DefaultListingType:       ConvertListingTypeToIndex017(in.DefaultListingType),
		InterfaceLanguage:        in.InterfaceLanguage,
		ShowAvatars:              in.ShowAvatars,
		SendNotificationsToEmail: in.SendNotificationsToEmail,
		ValidatorTime:            "1970-01-01T00:00:00Z",
		ShowBotAccounts:          in.ShowBotAccounts,
		ShowScores:               in.ShowScores,
		ShowReadPosts:            in.ShowReadPosts,
		ShowNewPostNotifs:        true,
		EmailVerified:            in.EmailVerified,
		AcceptedApplication:      in.AcceptedApplication,
	}
}

func ConvertLocalUserSettingsViewTo017(in lemmy.LocalUserView) lemmy017.LocalUserSettingsView {
	return lemmy017.LocalUserSettingsView{
		LocalUser: ConvertLocalUserSettingsTo017(in.LocalUser),
		Person:    ConvertPersonTo017(in.Person),
		Counts:    ConvertPersonAggregatesTo017(in.Counts),
	}
}

// ConvertMyUserInfoTo017 assembles a full 0.17.x-shaped MyUserInfo from
// the canonical one. InstanceBlocks is dropped — 0.17.2 has no such
// concept.
func ConvertMyUserInfoTo017(in *lemmy.MyUserInfo) *lemmy017.MyUserInfo {
	if in == nil {
		return nil
	}

	return &lemmy017.MyUserInfo{
		LocalUserView:       ConvertLocalUserSettingsViewTo017(in.LocalUserView),
		Follows:             helper.MapSlice(in.Follows, ConvertCommunityFollowerViewTo017),
		Moderates:           helper.MapSlice(in.Moderates, ConvertCommunityModeratorViewTo017),
		CommunityBlocks:     helper.MapSlice(in.CommunityBlocks, ConvertCommunityBlockViewTo017),
		PersonBlocks:        helper.MapSlice(in.PersonBlocks, ConvertPersonBlockViewTo017),
		DiscussionLanguages: in.DiscussionLanguages,
	}
}
// sortIndexTo017Type is the inverse of sortTypeTo017Index, used when
// parsing an incoming 0.17.x request (numeric index) into canonical's
// string-based SortType.
var sortIndexTo017Type = map[int16]lemmy.SortType{
	0:  lemmy.SortTypeActive,
	1:  lemmy.SortTypeHot,
	2:  lemmy.SortTypeNew,
	3:  lemmy.SortTypeOld,
	4:  lemmy.SortTypeTopDay,
	5:  lemmy.SortTypeTopWeek,
	6:  lemmy.SortTypeTopMonth,
	7:  lemmy.SortTypeTopYear,
	8:  lemmy.SortTypeTopAll,
	9:  lemmy.SortTypeMostComments,
	10: lemmy.SortTypeNewComments,
}

// ConvertIndex017ToSortType converts a 0.17.2 numeric sort-type index
// back to canonical's string-based SortType. Falls back to Hot for an
// out-of-range index.
func ConvertIndex017ToSortType(index int16) lemmy.SortType {
	if sortType, ok := sortIndexTo017Type[index]; ok {
		return sortType
	}
	return lemmy.SortTypeHot
}

// listingIndexTo017Type is the inverse of listingTypeTo017Index.
var listingIndexTo017Type = map[int16]lemmy.ListingType{
	0: lemmy.ListingTypeAll,
	1: lemmy.ListingTypeLocal,
	2: lemmy.ListingTypeSubscribed,
}

// ConvertIndex017ToListingType converts a 0.17.2 numeric listing-type
// index back to canonical's string-based ListingType. Falls back to
// All for an out-of-range index.
func ConvertIndex017ToListingType(index int16) lemmy.ListingType {
	if listingType, ok := listingIndexTo017Type[index]; ok {
		return listingType
	}
	return lemmy.ListingTypeAll
}

// ConvertPostViewTo017 assembles a full 0.17.x-shaped PostView from the
// canonical one, applying both conversions above.
func ConvertPostViewTo017(in lemmy.PostView) lemmy017.PostView {
	return lemmy017.PostView{
		Post:                       in.Post,
		Creator:                    ConvertPersonTo017(in.Creator),
		Community:                  in.Community,
		CreatorBannedFromCommunity: in.CreatorBannedFromCommunity,
		Counts:                     ConvertPostAggregatesTo017(in.Counts, in.Post.FeaturedCommunity, in.Post.FeaturedLocal),
		Subscribed:                 in.Subscribed,
		Saved:                      in.Saved,
		Read:                       in.Read,
		CreatorBlocked:             in.CreatorBlocked,
		MyVote:                     in.MyVote,
		UnreadComments:             in.UnreadComments,
	}
}
