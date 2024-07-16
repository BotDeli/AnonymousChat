package data

import (
	"fmt"
)

type TagID = string
type InterlocutorID = int64

type Interlocutor struct {
	Tag              TagID
	Source           SourceID
	ID               InterlocutorID
	SelfGender       GenderID
	TargetGender     GenderID
	CountConnections int
	SumDonation      int
}

func NewInterlocutorFromTelegram(id InterlocutorID, gender GenderID) Interlocutor {
	return Interlocutor{
		Tag:              formatTag(Telegram, id),
		Source:           Telegram,
		ID:               id,
		SelfGender:       gender,
		TargetGender:     Any,
		CountConnections: 0,
		SumDonation:      0,
	}
}

func formatTag(source SourceID, id InterlocutorID) TagID {
	return fmt.Sprintf("%s:%v", source, id)
}
