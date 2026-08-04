package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetPostsResponse struct {
	Posts    []piefed.PostView `json:"posts" validate:"required"`
	NextPage *string           `json:"next_page"`
}
