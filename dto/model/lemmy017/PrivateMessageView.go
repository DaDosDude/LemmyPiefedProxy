package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

type PrivateMessageView struct {
	PrivateMessage lemmy.PrivateMessage `json:"private_message" validate:"required"`
	Creator        Person               `json:"creator" validate:"required"`
	Recipient      Person               `json:"recipient" validate:"required"`
}
