package lemmy017

// MyUserInfo matches Lemmy 0.17.2's real MyUserInfo. InstanceBlocks
// (canonical has it) is dropped — 0.17.2 has no such concept.
type MyUserInfo struct {
	LocalUserView       LocalUserSettingsView    `json:"local_user_view" validate:"required"`
	Follows             []CommunityFollowerView  `json:"follows" validate:"required"`
	Moderates           []CommunityModeratorView `json:"moderates" validate:"required"`
	CommunityBlocks     []CommunityBlockView     `json:"community_blocks" validate:"required"`
	PersonBlocks        []PersonBlockView        `json:"person_blocks" validate:"required"`
	DiscussionLanguages []uint                   `json:"discussion_languages" validate:"required"`
}
