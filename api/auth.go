package api

import "time"

type MeResponse struct {
	CounsellorID   string    `json:"counsellorID"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Role           int       `json:"role"`
	Status         int       `json:"status"`
	Gender         int       `json:"gender"`
	Age            int       `json:"age"`
	IdentityNumber string    `json:"identityNumber"`
	PhoneNumber    string    `json:"phoneNumber"`
	LastLogin      time.Time `json:"lastLogin"`
	Avatar         string    `json:"avatar"`
	Email          string    `json:"email"`
	Title          string    `json:"title"`
	Department     string    `json:"department"`
	Qualification  string    `json:"qualification"`
	Introduction   string    `json:"introduction"`
	MaxConsults    int       `json:"maxConsults"`
}