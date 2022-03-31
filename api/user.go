package api

import "time"

type CounsellorInfoResponse struct {
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
	IsOnline       int       `json:"isOnline"`
}

type BindingInfoResponse struct {
	BindingID    int64  `json:"bindingID"`
	SupervisorID string `json:"supervisorID"`
	Name         string `json:"name"`
}

type SuperuserGetResponse struct {
	CounsellorID   string    `json:"counsellorID"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
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

type GetCounsellorListResponse struct {
	CounsellorList []*CounsellorInfoResponse `json:"counsellorList"`
	TotalCount     int                       `json:"totalCount"`
}
