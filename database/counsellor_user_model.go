package database

import "time"

type CounsellorUser struct {
	Username       string
	Name           string
	Password       string
	Role           int
	Status         int
	Gender         int
	Age            int
	IdentityNumber string
	PhoneNumber    string
	LastLogin      time.Time
	Avatar         string
	Email          string
	Title          string
	Department     string
	Qualification  string
	Introduction   string
	MaxConsults    int
}

func (CounsellorUser) TableName() string {
	return "counsellor_user"
}
