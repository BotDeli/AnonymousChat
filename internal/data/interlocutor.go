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
	CountConnections int
	SumDonation      int
}

func NewInterlocutorFromTelegram(id InterlocutorID) *Interlocutor {
	return &Interlocutor{
		Tag:              formatTag(Telegram, id),
		Source:           Telegram,
		ID:               id,
		CountConnections: 0,
		SumDonation:      0,
	}
}

func formatTag(source SourceID, id InterlocutorID) TagID {
	return fmt.Sprintf("%s:%v", source, id)
}
