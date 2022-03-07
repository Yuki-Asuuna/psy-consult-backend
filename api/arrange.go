package api

import "time"

type ArrangeResponse struct {
	ArrangeID    int64     `json:"arrangeID""`
	ArrangeTime  time.Time `json:"arrangeTime"`
	Role         int       `json:"role"`
	CounsellorID string    `json:"counsellorID"`
	Name         string    `json:"name"`
}

type GetArrangeResponse struct {
	Counsellors []*ArrangeResponse `json:"counsellors"`
	Supervisors []*ArrangeResponse `json:"supervisors"`
}
