package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetPrivateMessagesResponse struct {
	PrivateMessages []lemmy.PrivateMessageView `json:"private_messages" validate:"required"`
}
