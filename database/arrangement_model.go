package database

import "time"

type Arrangement struct {
	ArrangeID    int64
	ArrangeTime  time.Time
	Role         int
	CounsellorID string
}

func (Arrangement) TableName() string {
	return "arrangement"
}
