package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// CommunityResponse matches Lemmy 0.17.2's real CommunityResponse,
// returned by follow_community. Own fields match canonical exactly;
// only the nested CommunityView needs 0.17-shaping.
type CommunityResponse struct {
	CommunityView       model017.CommunityView `json:"community_view" validate:"required"`
	DiscussionLanguages []uint                 `json:"discussion_languages" validate:"required"`
}
