package lemmy

type PrivateMessageView struct {
	PrivateMessage PrivateMessage `json:"private_message" validate:"required"`
	Creator        Person         `json:"creator" validate:"required"`
	Recipient      Person         `json:"recipient" validate:"required"`
}
