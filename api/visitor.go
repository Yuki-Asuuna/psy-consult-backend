package api

import "time"

type VisitorInfoResponse struct {
	VisitorID        string    `json:"visitorID"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Status           int       `json:"status"`
	Gender           int       `json:"gender"`
	PhoneNumber      string    `json:"phoneNumber"`
	LastLogin        time.Time `json:"lastLogin"`
	Email            string    `json:"email"`
	EmergencyContact string    `json:"emergencyContact"`
	EmergencyPhone   string    `json:"emergencyPhone"`
	HasAgreed        int       `json:"hasAgreed"`
}
