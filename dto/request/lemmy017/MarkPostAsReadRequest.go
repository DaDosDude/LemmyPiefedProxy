package lemmy017

// PostId is a single value here — confirmed against Lemmy 0.17.2's real
// MarkPostAsRead struct, which has no post_ids batch array at all (that's
// a later addition our canonical model supports but 0.17.x doesn't).
// Read is deliberately NOT validate:"required" — false is legitimate
// (marking unread), and go-playground/validator's "required" tag rejects
// the zero value on plain bool fields.
type MarkPostAsReadRequest struct {
	PostId uint `json:"post_id" validate:"required"`
	Read   bool `json:"read"`
}
