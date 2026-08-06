package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/lemmy017"
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
		RegistrationMode:           string(in.RegistrationMode),
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
