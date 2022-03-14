package api

import "time"

type ConversationSearchResponse struct {
	ConversationID int64                   `json:"conversationID"`
	StartTime      time.Time               `json:"startTime"`
	EndTime        time.Time               `json:"endTime"`
	Counsellor     *CounsellorInfoResponse `json:"counsellorInfo"`
	Visitor        *VisitorInfoResponse    `json:"visitorInfo"`
	Status         int64                   `json:"status"`
	IsHelped       int64                   `json:"isHelped"`
	Supervisor     *CounsellorInfoResponse `json:"supervisorInfo"`
	Evaluation     *EvaluationInfoResponse `json:"evaluation"`
}
