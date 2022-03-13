package database

import "time"

type Conversation struct {
	ConversationID     int64
	StartTime          time.Time
	EndTime            time.Time
	CounsellorID       string
	VisitorID          string
	Status             int64
	IsHelped           int64
	HelpedSupervisorID string
}

func (Conversation) TableName() string {
	return "conversation"
}
