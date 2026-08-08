package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type GetPrivateMessagesResponse struct {
	PrivateMessages []model017.PrivateMessageView `json:"private_messages" validate:"required"`
}
