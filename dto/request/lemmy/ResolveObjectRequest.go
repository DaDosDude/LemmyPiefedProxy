package lemmy

type ResolveObjectRequest struct {
	Q string `json:"q" validate:"required"`
}
