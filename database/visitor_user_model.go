package database

import "time"

type VisitorUser struct {
	VisitorID        string
	Username         string
	Name             string
	Status           int
	Gender           int
	PhoneNumber      string
	LastLogin        time.Time
	Email            string
	EmergencyContact string
	EmergencyPhone   string
	HasAgreed        int
}

func (VisitorUser) TableName() string {
	return "visitor_user"
}
