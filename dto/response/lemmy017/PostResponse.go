package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// PostResponse is what Lemmy 0.17.2 actually returns for post creation,
// editing, liking, and marking as read — confirmed against each of
// those four handlers in Lemmy's own source. Just post_view, none of the
// extra context GetPostResponse carries.
type PostResponse struct {
	PostView model017.PostView `json:"post_view" required:"true"`
}
