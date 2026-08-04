package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// CommentResponse is what Lemmy 0.17.2 actually returns for fetching,
// creating, editing, and liking a comment — confirmed against Lemmy's
// own source. Unlike Post, there's no separate lean/full split; all four
// operations return this same shape.
//
// FormId is a known, documented gap: real 0.17.x echoes back whatever
// form_id the client sent on create (used for client-side optimistic-UI
// correlation). This proxy doesn't currently thread that value through
// from request to response, so it's always nil here — most clients
// treat form_id as best-effort, not core functionality, but this is a
// real limitation, not a resolved one.
type CommentResponse struct {
	CommentView  model017.CommentView `json:"comment_view" validate:"required"`
	RecipientIds []uint               `json:"recipient_ids" validate:"required"`
	FormId       *string              `json:"form_id,omitempty"`
}
