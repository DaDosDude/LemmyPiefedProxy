package lemmy

type GetCommunityRequest struct {
	Id   *uint   `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
