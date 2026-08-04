package lemmy

// LemmyApiError is this client's own error type for unsuccessful
// responses from a real Lemmy instance. It's deliberately separate from
// dto/response/lemmy.ErrorResponse (used elsewhere as this proxy's own
// output shape to its callers) — that type has a field named Error, not
// a method, so it can't satisfy Go's error interface. This type unmarshals
// the same {"error": "..."} wire shape real Lemmy actually sends, while
// genuinely implementing error via ErrorCode() string.
type LemmyApiError struct {
	ErrorCode  string `json:"error"`
	StatusCode int    `json:"-"`
}

func (receiver *LemmyApiError) Error() string {
	return receiver.ErrorCode
}
