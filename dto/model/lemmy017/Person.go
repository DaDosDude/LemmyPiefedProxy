package lemmy017

// Person matches Lemmy 0.17.2's real PersonSafe struct. Two fields our
// canonical Person model doesn't have at all: InboxUrl and Admin — both
// required (non-Option) in 0.17.2's real struct, confirmed against
// Lemmy's own source. Missing required fields in a serde Deserialize
// target are a real, likely deserialization failure for a Rust client
// like lemmyBB, unlike extra unknown fields (which serde ignores by
// default) — this isn't a cosmetic gap.
//
// InboxUrl has no equivalent data in our canonical model at all; it's
// derived as actor_id + "/inbox", the standard ActivityPub convention
// (matches what real Lemmy/Piefed actually put there in practice).
// Admin has no equivalent data available at this specific embedding
// point (Person, not PersonView, has no admin flag in either version's
// model) — defaulted to false rather than guessed, a known, documented
// limitation, not a resolved one.
// BotAccount is a required plain bool here, not the optional pointer our
// canonical model uses — defaulted to false if unset there.
type Person struct {
	Id           uint    `json:"id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	DisplayName  *string `json:"display_name,omitempty"`
	Avatar       *string `json:"avatar,omitempty"`
	Banned       bool    `json:"banned" validate:"required"`
	Published    string  `json:"published" validate:"required"`
	Updated      *string `json:"updated,omitempty"`
	ActorId      string  `json:"actor_id" validate:"required"`
	Bio          *string `json:"bio,omitempty"`
	Local        bool    `json:"local" validate:"required"`
	Banner       *string `json:"banner,omitempty"`
	Deleted      bool    `json:"deleted" validate:"required"`
	InboxUrl     string  `json:"inbox_url" validate:"required"`
	MatrixUserId *string `json:"matrix_user_id,omitempty"`
	Admin        bool    `json:"admin"`
	BotAccount   bool    `json:"bot_account"`
	InstanceId   uint    `json:"instance_id" validate:"required"`
}
