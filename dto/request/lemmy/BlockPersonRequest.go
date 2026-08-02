package lemmy

// Block is deliberately NOT validate:"required" — false is a legitimate
// value here (unblocking), and go-playground/validator's "required" tag
// rejects the zero value on plain bool fields.
type BlockPersonRequest struct {
	PersonId uint `json:"person_id" validate:"required"`
	Block    bool `json:"block"`
}
